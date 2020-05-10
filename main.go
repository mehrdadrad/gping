package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/mehrdadrad/gping/proto"
	"github.com/mehrdadrad/ping"
)

type params struct {
	mode     bool
	host     string
	bind     string
	count    int
	ttl      int
	interval string
}

type server struct{}

func (s *server) GetPing(pingReq *pb.PingRequest, stream pb.Ping_GetPingServer) error {
	p, err := ping.New(pingReq.DstAddr)
	if err != nil {
		log.Fatal(err)
	}

	p.SetCount(int(pingReq.Count))
	p.SetTTL(int(pingReq.Ttl))
	p.SetInterval(pingReq.Interval)
	p.SetPrivilegedICMP(false)

	rc, err := p.RunWithContext(stream.Context())
	if err != nil {
		log.Fatal(err)
	}

	for r := range rc {
		stream.Send(&pb.PingReply{
			Rtt:  r.RTT,
			Ttl:  int32(r.TTL),
			Seq:  int32(r.Sequence),
			Addr: r.Addr,
		})
	}
	return nil
}

func main() {
	p, err := getCli()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
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
		&pb.PingRequest{
			DstAddr:  p.host,
			Count:    int32(p.count),
			Ttl:      int32(p.ttl),
			Interval: p.interval,
		})

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

		fmt.Println(fmtPingLine(p))
	}
}

func fmtPingLine(p *pb.PingReply) string {
	return fmt.Sprintf("64 bytes from %s: icmp_seq=%d ttl=%d time=%.3f ms",
		p.Addr, p.Seq, p.Ttl, p.Rtt)
}
