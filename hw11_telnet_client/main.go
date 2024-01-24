package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var defaultTimeoutFlag time.Duration

func init() {
	flag.DurationVar(&defaultTimeoutFlag, "timeout", 10*time.Second, "This default timeout")
}

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("specify the host and port")
	}

	args := flag.Args()
	host := args[0]
	port := args[1]

	addr := net.JoinHostPort(host, port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer l.Close()

	tClient := NewTelnetClient(addr, defaultTimeoutFlag, os.Stdin, os.Stdout)

	err = tClient.Connect()
	if err != nil {
		log.Panicf("Cannot tClient.Connect: %v", err)
	}
	defer tClient.Close()

	res := make(chan error)

	go func(res chan error) {
		res <- tClient.Receive()
	}(res)

	go func(res chan error) {
		res <- tClient.Send()
	}(res)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	defer signal.Stop(sigint)

	select {
	case <-sigint:
	case <-res:
		close(sigint)
	}
}
