package main

import "testing"

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
