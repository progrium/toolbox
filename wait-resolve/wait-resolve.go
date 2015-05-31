package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var timeout = flag.Duration("t", 0, "Duration before timeout")

func main() {
	if len(os.Args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: wait-resolve [-t=timeout] <hostname>")
		os.Exit(1)
	}

	flag.Parse()

	if *timeout != 0 {
		go func() {
			time.Sleep(*timeout)
			os.Exit(1)
		}()
	}

	for {
		_, err := net.LookupHost(flag.Arg(0))
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

}
