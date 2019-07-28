package main

import (
	"log"
	"net"

	"github.com/jacobkaufmann/gocoin/pkg/p2p"
)

const (
	// LocalHost is a convenience variable for the string representation of
	// the localhost IP address.
	LocalHost = "127.0.0.1"
)

var (
	// LocalHostIP is a convenience variable for the IP of the localhost.
	LocalHostIP = [4]byte{127, 0, 0, 1}
)

// Client represents a Bitcoin client.
type Client struct {
	Port        int
	ConnManager *p2p.ConnManager
}

// NewClient returns a new client at localhost:port.
func NewClient(port int) *Client {
	return &Client{
		Port:        port,
		ConnManager: p2p.NewConnManager(),
	}
}

// LocalAddr returns the local TCP address of the client.
func (c *Client) LocalAddr() *net.TCPAddr {
	return &net.TCPAddr{
		IP:   LocalHostIP[:],
		Port: c.Port,
	}
}

// Listen initializes the peer on its port and begins listening for inbound
// connections.
func (c *Client) Listen() {
	localAddr := c.LocalAddr()

	l, err := net.ListenTCP(p2p.NetTCP, localAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("listening at %v", l.Addr())

	for {
		// Wait for connections.
		conn, err := l.AcceptTCP()
		if err != nil {
			log.Printf("failed to establish connection with %v: %v", conn, err)
		}

		// Handle the connection.
		go c.handleConn(conn)
	}
}

// handleConn establishes a connection with a new peer at conn and adds the
// peer in the client's connection manager.
func (c *Client) handleConn(conn *net.TCPConn) {
	peer := p2p.NewPeer(conn, true)
	added := c.ConnManager.AddConn(peer)
	if added {
		log.Printf("connection established with %v at time: %v",
			peer.Conn.RemoteAddr(), peer.TimeConnected.UTC())
	}
}

// Dial attempts to connect to a peer at address. If successful, a peer is
// created and a connection is added to the connection manager.
func (c *Client) Dial(address string) error {
	log.Printf("attempting to dial peer at %v", address)
	conn, err := net.Dial(p2p.NetTCP, address)
	if err != nil {
		return err
	}
	log.Printf("successfully dialed peer")
	peer := p2p.NewPeer(conn.(*net.TCPConn), false)
	added := c.ConnManager.AddConn(peer)
	if added {
		log.Printf("connection established with %v at time: %v",
			peer.Conn.RemoteAddr(), peer.TimeConnected)
	}
	return nil
}
