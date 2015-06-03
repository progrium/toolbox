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
		fmt.Fprintln(os.Stderr, "Usage: wait-conn [-t=timeout] <address>")
		os.Exit(1)
	}

	flag.Parse()

	addr, err := net.ResolveTCPAddr("tcp", flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid address:", err)
		os.Exit(1)
	}

	if *timeout != 0 {
		go func() {
			time.Sleep(*timeout)
			os.Exit(1)
		}()
	}

	for {
		conn, err := net.DialTCP("tcp", nil, addr)
		if err == nil {
			conn.Close()
			break
		}
		time.Sleep(1 * time.Second)
	}

}
