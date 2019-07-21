package p2p

import "net"

const (
	// MaxOutboundConns is the maximum number of outgoing peer connections.
	MaxOutboundConns = 8

	// DefaultMaxPeerConns is the maximum number of peer connections to
	// maintain.
	DefaultMaxPeerConns = 125

	// PingInterval is the time between pings automatically sent out for
	// latency probing and keepalive.
	PingInterval = 2 * 60

	// TimeoutInterval is the time after which to disconnect from inactivity.
	TimeoutInterval = 20 * 60
)

// ConnManager manages the connections of a Bitcoin node.
type ConnManager struct {
	Conns map[string]*net.TCPConn
}

// NewConnManager creates a new connection manager.
func NewConnManager() *ConnManager {
	return &ConnManager{
		Conns: make(map[string]*net.TCPConn),
	}
}

// NumConns returns the number of connections under management of mgr.
func (mgr *ConnManager) NumConns() int {
	return len(mgr.Conns)
}
