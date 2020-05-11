# gping
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
64 bytes from 172.217.14.78: icmp_seq=0 ttl=54 time=4.129 ms
64 bytes from 172.217.14.78: icmp_seq=1 ttl=54 time=2.134 ms
64 bytes from 172.217.14.78: icmp_seq=2 ttl=54 time=2.220 ms
64 bytes from 172.217.14.78: icmp_seq=3 ttl=54 time=2.194 ms
64 bytes from 172.217.14.78: icmp_seq=4 ttl=54 time=5.920 ms

5 packets transmitted, 5 packets received, 0.0% packet loss
Round-trip min/avg/max = 2.134/4.177/5.920 ms
```
# Client side JSON format
```
gping -c 5 -json -remote 192.168.10.15:3055 googl.com
```
```json
{"Min":2.142,"Avg":20.54875,"Max":32.684,"PacketLoss":0,"PacketTransmitted":5}
```

# Quick Help
```
  gRPC Ping v0.2.0

  usage:
  gping [options] host  
  
  options:
  --server, -s                -server or -s runs server (default: false) [$GPING_SERVER]
  --count value, -c value     -count or -c sets ping count (default: 4) [$GPING_COUNT]
  --ttl value, -t value       -ttl or -t sets time to live outgoing packets (default: 64) [$GPING_TTL]
  --size value                -size bytes (data + ICMP header) (default: 64) [$GPING_SIZE]
  --interval value, -i value  -intervale or -i sets time interval in format ns,us,ms,s (default: "1s") [$GPING_INTERVAL]
  --remote value, -r value    -remote 192.168.10.12:3055 [$GPING_REMOTE]
  --bind value, -b value      -bind 192.168.10.12:3055 (default: "0.0.0.0:3055") [$GPING_BIND]
  --json                      -json prints statistics in json format (default: false) [$GPING_JSON]
  --silent                    -silent prints just statistics (default: false) [$GPING_SILENT]
  --help, -h                  show help (default: false)
  --version, -v               print the version (default: false)
```
