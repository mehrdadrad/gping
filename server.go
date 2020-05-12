package main

import (
	"log"
	"net"

	pb "github.com/mehrdadrad/gping/proto"
	"github.com/mehrdadrad/ping"
	"google.golang.org/grpc"
)

type server struct {
	privileged bool
}

func (s *server) GetPing(pingReq *pb.PingRequest, stream pb.Ping_GetPingServer) error {
	var errStr string

	p, err := ping.New(pingReq.DstAddr)
	if err != nil {
		return stream.Send(&pb.PingReply{Err: err.Error()})
	}

	p.SetCount(int(pingReq.Count))
	p.SetTTL(int(pingReq.Ttl))
	p.SetTOS(int(pingReq.Tos))
	p.SetPacketSize(int(pingReq.Size))
	p.SetSrcIPAddr(pingReq.SrcAddr)
	p.SetInterval(pingReq.Interval)
	p.SetTimeout(pingReq.Timeout)
	p.SetPrivilegedICMP(s.privileged)

	rc, err := p.RunWithContext(stream.Context())
	if err != nil {
		return stream.Send(&pb.PingReply{Err: err.Error()})
	}

	for r := range rc {
		if r.Err != nil {
			errStr = r.Err.Error()
		}
		stream.Send(&pb.PingReply{
			Rtt:  r.RTT,
			Ttl:  int32(r.TTL),
			Seq:  int32(r.Seq),
			Addr: r.Addr,
			Size: int32(r.Size),
			Err:  errStr,
		})

		errStr = ""
	}
	return nil
}

func pingServer(p params) *grpc.Server {
	l, err := net.Listen("tcp", p.bind)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("gping server has been started on ", l.Addr().String())

	s := grpc.NewServer()
	go func() {
		pb.RegisterPingServer(s, &server{p.privileged})
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	return s
}
