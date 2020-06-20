package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/mehrdadrad/gping/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func pingClient(ctx context.Context, p params) chan *pb.PingReply {
	resp := make(chan *pb.PingReply, 1)
	ops := []grpc.DialOption{}

	if p.cert != "" && p.key != "" {
		if creds, err := transportClientCreds(p.cert, p.key, p.caCert); err != nil {
			log.Fatal(err)
		} else {
			ops = append(ops, grpc.WithTransportCredentials(creds))
		}
	} else {
		ops = append(ops, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(p.remote, ops...)
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
			Hosts:    p.hosts,
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

func transportClientCreds(certFile, keyFile, caCertFile string) (credentials.TransportCredentials, error) {
	var caCertPool *x509.CertPool
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	if caCertFile != "" {
		caCert, err := ioutil.ReadFile(caCertFile)
		if err != nil {
			return nil, err
		}

		caCertPool = x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
	}

	tc := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		// if RootCAs is nil, TLS uses the host's root CA set
		RootCAs: caCertPool,
	})

	return tc, nil
}
