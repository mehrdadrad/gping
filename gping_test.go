package main

import (
	"os"
	"testing"
)

func TestServer(t *testing.T) {
	p := params{
		bind:     "127.0.0.1:3055",
		remote:   "127.0.0.1:3055",
		host:     "127.0.0.1",
		count:    1,
		interval: "1s",
		ttl:      64,
		size:     64,
	}

	s := pingServer(p)
	defer s.Stop()

	rc := pingClient(p)
	r := <-rc
	if r.Err != "" {
		t.Error(r.Err)
	}

	if r.Ttl != 64 {
		t.Error("expected TTL 64 but got", r.Ttl)
	}

	if r.Seq != 0 {
		t.Error("expected Seq 0 but got", r.Seq)
	}

	if r.Rtt == 0 {
		t.Error("expected RTT greater than zero but got", r.Rtt)
	}

	if r.Addr != "127.0.0.1" {
		t.Error("expected addr: 127.0.0.1 but got", r.Addr)
	}
}

func TestServerTimeout(t *testing.T) {
	p := params{
		bind:     "127.0.0.1:3055",
		remote:   "127.0.0.1:3055",
		host:     "127.0.0.55",
		count:    1,
		interval: "1s",
		ttl:      64,
		size:     64,
	}

	s := pingServer(p)
	defer s.Stop()

	rc := pingClient(p)
	r := <-rc
	if r.Err != "Request timeout" {
		t.Error(r.Err)
	}

}

func TestGetCLI(t *testing.T) {
	os.Args = []string{"gping", "-c", "100", "-remote", "localhost:3055", "google.com"}
	p, err := getCli()
	if err != nil {
		t.Error(err)
	}

	if p.count != 100 {
		t.Error("expected count 100 but got", p.count)
	}

	if p.host != "google.com" {
		t.Error("expected host google.com but got", p.host)
	}

	if p.remote != "localhost:3055" {
		t.Error("expected remote localhost:3055 but got", p.remote)
	}
}
