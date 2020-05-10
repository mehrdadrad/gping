# gping
[![Build Status](https://travis-ci.org/mehrdadrad/gping.svg?branch=master)](https://travis-ci.org/mehrdadrad/gping) 
[![Go Report Card](https://goreportcard.com/badge/github.com/mehrdadrad/gping)](https://goreportcard.com/report/github.com/mehrdadrad/gping)

gping is a network tool to ping a target from a remote host. it works as client-server arch through gRPC protocol. it doesn't execute the ping shell command at the remote host instead it runs ping through a [Golang ping library](https://github.com/mehrdadrad/ping). use cases can be measurement full mesh network latency between nodes / data center or ping targets from different data center without SSH access to remote host.

![gping](/gping.png?raw=true "gping")

# Server side
```
#gping -server
```

# Client side
```
#gping -c 5 -remote 192.168.10.15:3055 googl.com
64 bytes from 172.217.14.78: icmp_seq=0 ttl=54 time=1.977 ms
64 bytes from 172.217.14.78: icmp_seq=1 ttl=54 time=2.978 ms
64 bytes from 172.217.14.78: icmp_seq=2 ttl=54 time=2.233 ms
64 bytes from 172.217.14.78: icmp_seq=3 ttl=54 time=2.395 ms
```
# Quick Help
```
  gRPC Ping v0.1.0

  usage:
  gping [options] host  
  
  options:
  --server, -s                -server or -s runs server (default: false) [$GPING_SERVER]
  --count value, -c value     -count or -c (default: 4) [$GPING_COUNT]
  --ttl value, -t value       -ttl or -t (default: 64) [$GPING_TTL]
  --size value                -size bytes (data + ICMP header) (default: 64) [$GPING_SIZE]
  --interval value, -i value  -intervale or -i 2s (default: "1s") [$GPING_INTERVAL]
  --remote value, -r value    -remote 192.168.10.12:3055 [$GPING_REMOTE]
  --bind value, -b value      -bind 192.168.10.12:3055 (default: "0.0.0.0:3055") [$GPING_BIND]
  --help, -h                  show help (default: false)
  --version, -v               print the version (default: false)
```
