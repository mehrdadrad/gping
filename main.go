package main

import (
	"fmt"
	"os"
)

type params struct {
	mode     bool
	host     string
	bind     string
	count    int
	ttl      int
	interval string
}

func main() {
	p, err := getCli()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if p.mode {
		pingServer(p)
	} else {
		pingClient(p)
	}

}
