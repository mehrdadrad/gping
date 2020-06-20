package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"github.com/mehrdadrad/gping/proto"
	pb "github.com/mehrdadrad/gping/proto"
	"github.com/mehrdadrad/ping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

var isLogs bool

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

func (s *server) GetBulkPing(ctx context.Context, pingBulkReq *pb.PingBulkRequest) (*pb.PingBulkResult, error) {
	var (
		wg          sync.WaitGroup
		pingResChan = make(chan *pb.PingResult, 1)
	)

	for _, host := range pingBulkReq.Hosts {
		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			r := s.pingWithResult(ctx, host, pingBulkReq)
			pingResChan <- r
		}(host)
	}

	go func() {
		wg.Wait()
		close(pingResChan)
	}()

	results := &pb.PingBulkResult{}
	for r := range pingResChan {
		results.Results = append(results.Results, r)
	}

	return results, nil
}

func (s *server) pingWithResult(ctx context.Context, host string, req *pb.PingBulkRequest) *pb.PingResult {
	pr := &pb.PingResult{}
	pr.Host = host
	p, err := ping.New(host)
	if err != nil {
		pr.Err = err.Error()
		return pr
	}

	p.SetCount(int(req.Count))
	p.SetTTL(int(req.Ttl))
	p.SetTOS(int(req.Tos))
	p.SetPacketSize(int(req.Size))
	p.SetSrcIPAddr(req.SrcAddr)
	p.SetInterval(req.Interval)
	p.SetTimeout(req.Timeout)
	p.SetPrivilegedICMP(s.privileged)

	rc, err := p.RunWithContext(ctx)
	if err != nil {
		pr.Err = err.Error()
		return pr
	}

	for r := range rc {
		if r.Err != nil {
			pr.PacketLoss++
			pr.Err = r.Err.Error()
			continue
		}
		pr.MinRtt = min(pr.MinRtt, r.RTT)
		pr.AvgRtt = avg(pr.AvgRtt, r.RTT)
		pr.MaxRtt = max(pr.MaxRtt, r.RTT)
	}

	pr.PacketLoss = (pr.PacketLoss * 100) / req.Count

	return pr
}

func pingServer(p params) *grpc.Server {
	isLogs = p.isLogs

	l, err := net.Listen("tcp", p.bind)
	if err != nil {
		log.Fatal(err)
	}

	if isLogs {
		log.Println("gping server has been started on ", l.Addr().String())
	}

	ops := []grpc.ServerOption{}
	ops = append(ops, grpc.StreamInterceptor(gpingInterceptor))

	if p.cert != "" && p.key != "" && p.clientsCert != "" {
		if creds, err := transportCreds(p.cert, p.key, p.clientsCert); err != nil {
			log.Fatal(err)
		} else {
			ops = append(ops, grpc.Creds(creds))
			if isLogs {
				log.Println("mTLS has been enabled")
			}
		}
	}

	s := grpc.NewServer(ops...)

	go func() {
		pb.RegisterPingServer(s, &server{p.privileged})
		if err := s.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	return s
}

func gpingInterceptor(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	var clientAddr string
	if p, ok := peer.FromContext(ss.Context()); ok && isLogs {
		clientAddr = p.Addr.String()
	}
	err := handler(srv, &wrappedStream{ss, clientAddr})
	return err
}

type wrappedStream struct {
	grpc.ServerStream
	clientAddr string
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)
	if pingReq, ok := m.(*proto.PingRequest); ok && isLogs && err == nil {
		log.Printf("received a ping request from %s to %s", w.clientAddr, pingReq.DstAddr)
	}
	return err
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	return w.ServerStream.SendMsg(m)
}

func transportCreds(certFile, keyFile, clientsCertFile string) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	clientsCert, err := ioutil.ReadFile(clientsCertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(clientsCert)

	tc := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	return tc, nil
}
