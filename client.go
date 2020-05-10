package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/mehrdadrad/gping/proto"
	"google.golang.org/grpc"
)

func pingClient(p params) {
	conn, err := grpc.Dial(p.remote, grpc.WithInsecure())
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
			Size:     int32(p.size),
			Interval: p.interval,
		})

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < p.count; i++ {
		p, err := r.Recv()
		if err != nil {
			if err == io.EOF {
				os.Exit(1)
			} else {
				log.Fatal(err)
			}
		}

		fmt.Println(fmtPingLine(p))
	}
}

func fmtPingLine(p *pb.PingReply) string {
	if p.Err != "" {
		return fmt.Sprintf("error: %s", p.Err)
	} else {
		return fmt.Sprintf("%d bytes from %s: icmp_seq=%d ttl=%d time=%.3f ms",
			p.Size, p.Addr, p.Seq, p.Ttl, p.Rtt)
	}
}
