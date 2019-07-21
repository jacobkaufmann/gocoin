package p2p

import (
	"log"
	"net"
	"time"

	"github.com/jacobkaufmann/gocoin/pkg/protocol"
)

const (
	// NetTCP is a convenience variable for managing TCP connections.
	NetTCP = "tcp"

	// MainnetPort is the default port for the Bitcoin network.
	MainnetPort = 8333

	// TestnetPort is the default port for the test Bitcoin network.
	TestnetPort = 18333

	// MaxSendBufferSize is the maximum number of messages to hold in a peer's
	// send message buffer.
	MaxSendBufferSize = 1000

	// MaxReceiveBufferSize is the maximum number of messages to hold in a peer's
	// receive message buffer.
	MaxReceiveBufferSize = 5000
)

// Peer represents a node on the Bitcoin network.
type Peer struct {
	// Conn is the TCP connection to the peer.
	Conn *net.TCPConn

	// Services is the service flag of the peer.
	Services protocol.ServiceFlag

	// Version is the protocol version of the peer.
	Version uint32

	// Net is the network of the peer.
	Net protocol.BitcoinNet

	// Inbound is a flag denoting whether the peer is an inbound connection.
	Inbound bool

	// TimeConnected is the time the connection to the peer was established.
	TimeConnected time.Time

	sendMsgBuf chan protocol.Message
	recvMsgBuf chan protocol.Message
}

// NewPeer returns a new peer.
func NewPeer(conn *net.TCPConn, inbound bool) *Peer {
	return &Peer{
		Conn:       conn,
		Services:   0,
		Inbound:    inbound,
		sendMsgBuf: make(chan protocol.Message, MaxSendBufferSize),
		recvMsgBuf: make(chan protocol.Message, MaxReceiveBufferSize),
	}
}

// EnqueueSendMessage enqueues a message to the peer's send message buffer.
func (p *Peer) EnqueueSendMessage(msg protocol.Message) {
	p.sendMsgBuf <- msg
}

// DequeueSendMessage attempts to dequeue a message from the peer's send message
// buffer.
func (p *Peer) DequeueSendMessage() protocol.Message {
	select {
	case msg := <-p.sendMsgBuf:
		return msg
	default:
		log.Println("no messages available to send.")
		return nil
	}
}

// EnqueueReceiveMessage enqueues a message to the peer's receive message buffer.
func (p *Peer) EnqueueReceiveMessage(msg protocol.Message) {
	p.recvMsgBuf <- msg
}

// DequeueReceiveMessage attempts to receive a message from the peer's receive
// message buffer.
func (p *Peer) DequeueReceiveMessage() protocol.Message {
	select {
	case msg := <-p.recvMsgBuf:
		return msg
	default:
		log.Println("no messages available to recieve.")
		return nil
	}
}

// SendMessage sends a message to p over its TCP connection.
func SendMessage(msg protocol.Message, p *Peer) (int, error) {
	return protocol.WriteMessage(p.Conn, msg, p.Version, p.Net)
}
