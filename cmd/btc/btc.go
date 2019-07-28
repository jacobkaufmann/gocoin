package main

import (
	"flag"
	"log"
	"net"

	"github.com/jacobkaufmann/gocoin/pkg/p2p"
)

var (
	port    int
	connect string
)

func init() {
	flag.IntVar(&port, "port", p2p.MainnetPort, "port on which to run")
	flag.StringVar(&connect, "connect", "", "ip:port of initial peer")
}

func main() {
	flag.Parse()
	client := NewClient(port)
	if connect != "" {
		log.Printf("attempting to connect to peer at %v", connect)
		conn, err := net.Dial(p2p.NetTCP, connect)
		if err != nil {
			log.Fatalf("failed to connect to peer at %v: %v", connect, err)
		}
		log.Printf("successfully dialed peer at %v", conn.RemoteAddr())
	}
	client.Listen()
}
