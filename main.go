package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	var (
		pr  = printer{}
		sig = make(chan os.Signal, 1)
	)

	signal.Notify(sig, os.Interrupt)

	p, err := getCli()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if p.mode {
		s := pingServer(p)
		<-sig
		s.Stop()
	} else {
		for r := range pingClient(p) {
			pr.print(r, p, sig)
		}
	}
}
