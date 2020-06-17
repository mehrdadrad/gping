# gping
[![Go Report Card](https://goreportcard.com/badge/github.com/mehrdadrad/gping)](https://goreportcard.com/report/github.com/mehrdadrad/gping)

gping is a network tool to ping a target from a remote host. it works as client-server arch through gRPC protocol. it doesn't execute the ping shell command at the remote host instead it runs ping through [Golang ping library](https://github.com/mehrdadrad/ping). use cases can be measurement full mesh network latency between nodes / data center or ping targets from different data center without SSH access to remote host.

![gping](/gping.png?raw=true "gping")

## Server side
```
#gping -server
```

## Single host
```
#gping -c 5 -r 192.168.10.15:3055 googl.com
64 bytes from 172.217.14.78: icmp_seq=0 ttl=54 time=4.129 ms
64 bytes from 172.217.14.78: icmp_seq=1 ttl=54 time=2.134 ms
64 bytes from 172.217.14.78: icmp_seq=2 ttl=54 time=2.220 ms
64 bytes from 172.217.14.78: icmp_seq=3 ttl=54 time=2.194 ms
64 bytes from 172.217.14.78: icmp_seq=4 ttl=54 time=5.920 ms

5 packets transmitted, 5 packets received, 0.0% packet loss
Round-trip min/avg/max = 2.134/4.177/5.920 ms
```
## Multiple hosts
```
gping -c 5 -r 192.168.10.15:3055 yahoo.com google.com juniper.com
```
```json
[
  {
    "addr": "google.com",
    "min_rtt": 2.531,
    "avg_rtt": 3.859,
    "max_rtt": 4.859
  },
  {
    "addr": "yahoo.com",
    "min_rtt": 2.337,
    "avg_rtt": 3.615,
    "max_rtt": 4.431
  },
  {
    "addr": "juniper.com",
    "packet_loss": 100,
    "err": "request timeout"
  }
]
```

## Single host JSON format
```
gping -c 5 -json -r 192.168.10.15:3055 yahoo.com
```
```json
{"min_rtt":2.444,"avg_rtt":3.855,"max_rtt":4.625}
```



## Quick Help
```
  gRPC Ping v0.2.0

  usage:
  gping [options] host(s) 
  
  options:
  --count value, -c value     sets ping count (default: 4) [$GPING_COUNT]
  --ttl value, -t value       sets the IP Time to Live (default: 64) [$GPING_TTL]
  --tos value, -q value       sets quality of service in ICMP datagram (default: 0) [$GPING_TOS]
  --size value                sets the number of data bytes to be sent (data + ICMP header) (default: 64) [$GPING_SIZE]
  --interval value, -i value  sets wait between sending each packet in format ns,us,ms,s (default: "1s") [$GPING_INTERVAL]
  --timeout value, -W value   sets time to wait for an ICMP reply in format ns,us,ms,s (default: "2s") [$GPING_TIMEOUT]
  --remote value, -r value    sets remote server IP_ADDR:PORT [$GPING_REMOTE]
  --json                      prints statistics in json format (default: false) [$GPING_JSON]
  --silent                    prints just statistics (default: false) [$GPING_SILENT]
  --server, -s                runs server (default: false) [$GPING_SERVER]
  --bind value, -b value      sets bind IP_ADDR:PORT [server] (default: "0.0.0.0:3055") [$GPING_BIND]
  --privileged, -p            enables ICMP privileged mode [server] (default: false) [$GPING_PRIVILEGED]
  --logs                      enables logging [server] (default: false) [$GPING_LOGS]
  --help, -h                  show help (default: false)
  --version, -v               print the version (default: false)
```

## Build
It can be built for Linux and macOS
```
#go build
```
note: server side works in two modes 1- unprivileged icmp (default) 2- privileged icmp

the first case you need to run the below command on Linux:
```
sudo sysctl -w net.ipv4.ping_group_range="0   2147483647"
```
in case of privileged icmp (it enables w/ -p or -privileged) you should give the raw socket permission by the below or just run by superuser.
```
sudo setcap cap_net_raw+ep ./gping
```
for more information about setcap cap_net_raw+ep: https://wiki.archlinux.org/index.php/Capabilities

## Supported platform

|                    | Client | Server |
|--------------------| -------|--------|
|Linux               | &check;| &check;| 
|macOS               | &check;| &check;| 
|FreeBSD             | &check;| &check;| 
|Windows             | &check;|        | 

## Download gping v0.2.0 
In case you don't want to compile, the below binaries are available (compiled by Go 1.14.2)
- [Linux](https://github.com/mehrdadrad/gping/releases/download/v0.2.0/gping-linux)
- [macOS](https://github.com/mehrdadrad/gping/releases/download/v0.2.0/gping-macos)
- [FreeBSD](https://github.com/mehrdadrad/gping/releases/download/v0.2.0/gping-freebsd)
- [Windows](https://github.com/mehrdadrad/gping/releases/download/v0.2.0/gping.exe)

## License
This project is licensed under MIT license. Please read the LICENSE file.


## Contribute
Welcomes any kind of contribution, please follow the next steps:

- Fork the project on github.com.
- Create a new branch.
- Commit changes to the new branch.
- Send a pull request.
