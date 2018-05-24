package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/m90/lclhst"
)

func main() {
	var (
		port    = flag.Int("port", 8080, "the port of the application")
		timeout = flag.Duration("timeout", 10*time.Second, "timeout for giving up")
	)
	flag.Parse()

	if err := lclhst.WaitDuration(*timeout, *port); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
