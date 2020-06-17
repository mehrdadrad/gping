package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/mehrdadrad/gping/proto"
)

type printer struct {
	min     float64
	avg     float64
	max     float64
	counter int
	loss    int
}

func (pr *printer) print(ctx context.Context, p params, ping *proto.PingReply) {
	pr.counter++

	if len(ping.Err) < 1 {
		pr.min = min(pr.min, ping.Rtt)
		pr.avg = avg(pr.avg, ping.Rtt)
		pr.max = max(pr.max, ping.Rtt)
	} else {
		pr.loss++
	}

	if p.json {
		select {
		case <-ctx.Done():
			pr.statisticsJSON()
			os.Exit(0)
		default:
			if p.count == pr.counter {
				pr.statisticsJSON()
			}
		}
		return
	}

	if p.count == pr.counter {
		printLine(ping, p)
		pr.statistics()
	} else {
		select {
		case <-ctx.Done():
			printLine(ping, p)
			pr.statistics()
			os.Exit(0)
		default:
			printLine(ping, p)
		}
	}

}

func (pr *printer) statistics() {
	fmt.Printf("\n%d packets transmitted, %d packets received, %.1f%% packet loss\n",
		pr.counter, pr.counter-pr.loss, float64(pr.loss*100)/float64(pr.counter))
	fmt.Printf("Round-trip min/avg/max = %.3f/%.3f/%.3f ms\n",
		pr.min, pr.avg, pr.max)
}

func (pr *printer) statisticsJSON() {
	a := struct {
		Min               float64
		Avg               float64
		Max               float64
		PacketLoss        float64
		PacketTransmitted int
	}{
		pr.min,
		pr.avg,
		pr.max,
		float64(pr.loss*100) / float64(pr.counter),
		pr.counter,
	}

	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}

func printLine(ping *proto.PingReply, p params) {
	if p.silent {
		return
	}

	if ping.Err != "" {
		fmt.Printf("%s for icmp_seq %d\n", ping.Err, ping.Seq)
	} else {
		fmt.Printf("%d bytes from %s: icmp_seq=%d ttl=%d time=%.3f ms\n",
			ping.Size, ping.Addr, ping.Seq, ping.Ttl, ping.Rtt)
	}
}

func avg(a, b float64) float64 {
	if a == 0.0 {
		return b
	}
	return precision((a + b) / 2)
}

func min(a, b float64) float64 {
	if a == 0.0 {
		return b
	}

	if a < b {
		return a
	}

	return b
}

func max(a, b float64) float64 {
	if a == 0.0 {
		return b
	}

	if a < b {
		return b
	}

	return a
}

func precision(f float64) float64 {
	f, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", f), 64)
	return f
}
