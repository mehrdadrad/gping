package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	pb "github.com/mehrdadrad/gping/proto"
	"google.golang.org/grpc"
)

func pingClient(ctx context.Context, p params) chan *pb.PingReply {
	resp := make(chan *pb.PingReply, 1)

	conn, err := grpc.Dial(p.remote, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewPingClient(conn)
	r, err := c.GetPing(
		ctx,
		&pb.PingRequest{
			DstAddr:  p.hosts[0],
			Count:    int32(p.count),
			Ttl:      int32(p.ttl),
			Tos:      int32(p.tos),
			Size:     int32(p.size),
			Interval: p.interval,
			Timeout:  p.timeout,
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

func pingBulkClient(ctx context.Context, p params) {
	conn, err := grpc.Dial(p.remote, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := pb.NewPingClient(conn)
	resp, err := c.GetBulkPing(ctx,
		&pb.PingBulkRequest{
			DstAddrs: p.hosts,
			Count:    int32(p.count),
			Ttl:      int32(p.ttl),
			Tos:      int32(p.tos),
			Size:     int32(p.size),
			Interval: p.interval,
			Timeout:  p.timeout,
		})
	if err != nil {
		log.Fatal(err)
	}

	r := resp.GetResults()
	b, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(b)
}
