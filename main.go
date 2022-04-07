package main

import (
	"flag"
	"time"
)

var (
	timeout = flag.Int("t", 0, "timeout")
)

func main() {
	t := time.Millisecond * time.Duration(*timeout)

	time.Sleep(t)
}
