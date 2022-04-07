package main

import (
	"flag"
	"time"
)

var (
	timeout = flag.Int("t", -1, "timeout")
)

func main() {
	flag.Parse()

	if *timeout == -1 {
		flag.Usage()
	}

	t := time.Millisecond * time.Duration(*timeout)

	time.Sleep(t)
}
