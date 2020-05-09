package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/mehrdadrad/gping/proto"
	"github.com/mehrdadrad/ping"
	cli "github.com/urfave/cli/v2"
)

type params struct {
	mode  bool
	host  string
	bind  string
	count int
	ttl   int
}

type server struct{}

var log grpclog.LoggerV2

func (s *server) GetPing(pingReq *pb.PingRequest, stream pb.Ping_GetPingServer) error {
	p, err := ping.New(pingReq.DstAddr)
	if err != nil {
		log.Fatal(err)
	}

	p.SetCount(int(pingReq.Count))
	p.SetInterval("1s")
	p.SetPrivilegedICMP(false)

	rc, err := p.RunWithContext(stream.Context())
	if err != nil {
		log.Fatal(err)
	}

	for r := range rc {
		stream.Send(&pb.PingReply{Rtt: r.RTT})
	}
	return nil
}

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func main() {
	p := params{}

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "-server or -s",
			EnvVars: []string{"GPING_SERVER"},
		},
		&cli.IntFlag{
			Name:    "count",
			Aliases: []string{"c"},
			Usage:   "-count or -c",
			EnvVars: []string{"GPING_COUNT"},
		},
		&cli.IntFlag{
			Name:    "ttl",
			Aliases: []string{"t"},
			Usage:   "-ttl or -t",
			EnvVars: []string{"GPING_TTL"},
		},
		&cli.StringFlag{
			Name:    "remote",
			Aliases: []string{"r"},
			Usage:   "-remote 192.168.10.12:3055",
			EnvVars: []string{"GPING_REMOTE"},
		},
		&cli.StringFlag{
			Name:    "bind",
			Aliases: []string{"b"},
			Usage:   "-bind 192.168.10.12:3055",
			EnvVars: []string{"GPING_BIND"},
		},
	}
	app := &cli.App{
		Flags: flags,
		Action: func(c *cli.Context) error {
			p.mode = c.Bool("server")
			if c.Int("count") == 0 {
				p.count = 4
			} else {
				p.count = c.Int("count")
			}

			p.host = c.Args().Get(0)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if p.mode {
		pingServer(p)
	} else {
		pingClient(p)
	}

}

func pingServer(p params) {
	l, err := net.Listen("tcp", ":8055")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterPingServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func pingClient(p params) {
	conn, err := grpc.Dial("localhost:8055", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewPingClient(conn)
	r, err := c.GetPing(
		context.Background(),
		&pb.PingRequest{DstAddr: p.host, Count: int32(p.count)})

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < p.count; i++ {
		p, err := r.Recv()
		if err == io.EOF {
			log.Fatal("end of ping!")
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%.2f\n", p.Rtt)
	}
}
