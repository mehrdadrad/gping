package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/mehrdadrad/gping/proto"
	"google.golang.org/grpc"
)

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
