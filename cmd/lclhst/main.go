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
		port     = flag.Int("port", 8080, "the port of the application")
		deadline = flag.Duration("deadline", 10*time.Second, "deadline for giving up")
	)
	flag.Parse()

	if err := lclhst.WaitFor(*deadline, *port); err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
