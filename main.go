package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	var (
		pr          = printer{}
		sig         = make(chan os.Signal, 1)
		ctx, cancel = context.WithCancel(context.Background())
	)

	signal.Notify(sig, os.Interrupt)

	go func() {
		<-sig
		cancel()
	}()

	p, err := getCli()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if p.mode {
		s := pingServer(p)
		<-sig
		s.Stop()
	} else if len(p.hosts) == 1 {
		for r := range pingClient(ctx, p) {
			pr.print(ctx, p, r)
		}
	} else {
		pingBulkClient(ctx, p)
	}
}
