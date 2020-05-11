package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/mehrdadrad/gping/proto"
	"google.golang.org/grpc"
)

func pingClient(p params) chan *pb.PingReply {
	resp := make(chan *pb.PingReply, 1)

	conn, err := grpc.Dial(p.remote, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

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

	go func() {
		for i := 0; i < p.count; i++ {
			p, err := r.Recv()
			if err != nil {
				resp <- &pb.PingReply{Err: err.Error()}
				break
			}
			resp <- p
		}

		conn.Close()
		close(resp)
	}()

	return resp
}

func pingPrint(ping *pb.PingReply, p params) {
	if ping.Err != "" {
		fmt.Printf("error: %s\n", ping.Err)
	}

	fmt.Printf("%d bytes from %s: icmp_seq=%d ttl=%d time=%.3f ms\n",
		ping.Size, ping.Addr, ping.Seq, ping.Ttl, ping.Rtt)
}
