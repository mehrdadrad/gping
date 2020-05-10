package main

import (
	"log"
	"net"

	pb "github.com/mehrdadrad/gping/proto"
	"github.com/mehrdadrad/ping"
	"google.golang.org/grpc"
)

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
