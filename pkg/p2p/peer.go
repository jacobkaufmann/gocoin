package p2p

import (
	"log"
	"net"

	"github.com/jacobkaufmann/gocoin/pkg/protocol"
)

const (
	// NetTCP is a convenience variable for managing TCP connections.
	NetTCP = "tcp"

	// MainnetPort is the default port for the Bitcoin network.
	MainnetPort = 8333

	// TestnetPort is the default port for the test Bitcoin network.
	TestnetPort = 18333
)

// Peer represents a node on the Bitcoin network.
type Peer struct {
	Conn       *net.TCPConn
	Services   protocol.ServiceFlag
	Version    uint32
	Inbound    bool
	bufSendMsg chan protocol.Message
	bufRecvMsg chan protocol.Message
}

// NewPeer returns a new peer.
func NewPeer(conn *net.TCPConn, inbound bool) *Peer {
	return &Peer{
		Conn:       conn,
		Services:   0,
		Inbound:    inbound,
		bufSendMsg: make(chan protocol.Message),
		bufRecvMsg: make(chan protocol.Message),
	}
}

// SendMsg sends a message to p over its TCP connection.
func (p *Peer) SendMsg(msg protocol.Message) error {
	return msg.Serialize(p.Conn, p.Version)
}

// EnqueueSendMsg enqueues a message to the peer's send message buffer.
func (p *Peer) EnqueueSendMsg(msg protocol.Message) {
	p.bufSendMsg <- msg
}

// DequeueSendMsg attempts to dequeue a message from the peer's send message
// buffer.
func (p *Peer) DequeueSendMsg() protocol.Message {
	select {
	case msg := <-p.bufSendMsg:
		return msg
	default:
		log.Println("no messages available to send.")
		return nil
	}
}

// EnqueueRecvMsg enqueues a message to the peer's receive message buffer.
func (p *Peer) EnqueueRecvMsg(msg protocol.Message) {
	p.bufRecvMsg <- msg
}

// DequeueRecvMsg attempts to receive a message from the peer's receive
// message buffer.
func (p *Peer) DequeueRecvMsg() protocol.Message {
	select {
	case msg := <-p.bufRecvMsg:
		return msg
	default:
		log.Println("no messages available to recieve.")
		return nil
	}
}
