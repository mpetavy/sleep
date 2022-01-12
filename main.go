package main

import (
	"flag"
	"fmt"
	"github.com/mpetavy/common"
	"time"
)

var (
	timeout = flag.Int("t", 0, "timeout")
)

func init() {
	common.Init(false, "0.0.0", "", "", "2018", "test", "mpetavy", fmt.Sprintf("https://github.com/mpetavy/%s", common.Title()), common.APACHE, nil, nil, nil, run, 0)
}

func run() error {
	t := common.MillisecondToDuration(*timeout)

	common.Debug("sleep %v...", t)
	time.Sleep(t)
	common.Debug("continue")

	return nil
}

func main() {
	defer common.Done()

	common.Run(nil)
}
