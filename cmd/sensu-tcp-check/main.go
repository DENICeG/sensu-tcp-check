package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	log.SetOutput(os.Stderr)

	if len(os.Args) != 3 {
		println("usage: sensu-tcp-check <metricprefix> <host:port>")
		os.Exit(2)
	}

	timeBegin := time.Now()

	conn, err := net.DialTimeout("tcp", os.Args[2], 10*time.Second)
	if err != nil {
		log.Printf("ERROR: %s\n\n", err.Error())
		fmt.Printf("%s,check=tcp %s=%d,%s=%d %d\n",
			os.Args[1],
			"available", 0,
			"duration", 0,
			timeBegin.Unix())

		if conn != nil {
			conn.Close()
		}

		os.Exit(2)
	}

	responseTime := time.Since(timeBegin).Milliseconds()

	log.Printf("OK\n\n")
	fmt.Printf("%s,check=tcp %s=%d,%s=%d %d\n",
		os.Args[1],
		"available", 1,
		"duration", responseTime,
		timeBegin.Unix())

	_ = conn.Close()

	os.Exit(0)
}
