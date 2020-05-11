package main

import (
	"fmt"
	"os"
	"sync"
)

type params struct {
	mode     bool
	host     string
	bind     string
	remote   string
	count    int
	ttl      int
	size     int
	interval string
}

func main() {
	var wg sync.WaitGroup

	p, err := getCli()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if p.mode {
		wg.Add(1)
		pingServer(p)
		wg.Wait()
	} else {
		for r := range pingClient(p) {
			pingPrint(r, p)
		}
	}
}
