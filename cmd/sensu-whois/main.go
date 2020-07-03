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

	if len(os.Args) == 1 {
		println("usage: sensu-tcp-check <host>:<port>")
		os.Exit(2)
	}

	timeBegin := time.Now()

	conn, err := net.DialTimeout("tcp", os.Args[1], 10*time.Second)
	if err != nil {
		log.Println(err.Error())
		fmt.Printf("%s %d %d\n", "sensu.tcp.available", 0, timeBegin.Unix())
		fmt.Printf("%s %d %d\n", "sensu.tcp.duration", 0, timeBegin.Unix())

		if conn != nil {
			conn.Close()
		}

		os.Exit(2)
	}

	responseTime := time.Since(timeBegin).Milliseconds()

	fmt.Printf("%s %d %d\n", "sensu.tcp.available", 1, timeBegin.Unix())
	fmt.Printf("%s %d %d\n", "sensu.tcp.duration", responseTime, timeBegin.Unix())

	_ = conn.Close()

	os.Exit(0)
}
