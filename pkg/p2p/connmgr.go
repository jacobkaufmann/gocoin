package p2p

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/jacobkaufmann/gocoin/pkg/protocol"
)

const (
	// MaxOutboundConns is the maximum number of outgoing peer connections.
	MaxOutboundConns = 8

	// MaxInboundConns is the maximum number of incoming peer connections to
	// maintain.
	MaxInboundConns = 125

	// PingInterval is the time between pings automatically sent out for
	// latency probing and keepalive.
	PingInterval = 2 * 60

	// TimeoutInterval is the time after which to disconnect from inactivity.
	TimeoutInterval = 20 * 60

	// MaxRetries is the number of retries before no more attempts are made to
	// connect to a peer.
	MaxRetries = 3
)

// ConnInfo holds statistics about a connection to a peer.
type ConnInfo struct {
	Source             *net.TCPAddr
	LastSuccessfulConn time.Time
}

// ConnManager manages the connections of a Bitcoin node.
type ConnManager struct {
	// Conns is the connection pool being managed. The connnection pool is
	// protected by a read-write mutex and is thread-safe.
	Conns    map[string]*Peer
	muxConns sync.RWMutex

	// NumOutbound is the number of outbound connections in the connection
	// pool.
	NumOutbound int

	// NumInbound is the number of inbound connections in the connection pool.
	NumInbound int

	// Services is the service flag of the local node.
	Services protocol.ServiceFlag
}

// NewConnManager creates a new connection manager.
func NewConnManager() *ConnManager {
	return &ConnManager{
		Conns:       make(map[string]*Peer),
		NumOutbound: 0,
		NumInbound:  0,
		Services:    0,
	}
}

// NumConns returns the number of connections under management of mgr.
func (mgr *ConnManager) NumConns() int {
	return len(mgr.Conns)
}

// GetConn retrieves a peer connection to addr from mgr's connection pool if
// such a connection exists.
func (mgr *ConnManager) GetConn(addr string) *Peer {
	mgr.muxConns.RLock()
	defer mgr.muxConns.RUnlock()
	p, ok := mgr.Conns[addr]
	if ok {
		return p
	}
	return nil
}

// RemoveConn removes the connection to addr from mgr's connection pool if
// such a connection exists.
func (mgr *ConnManager) RemoveConn(addr string) (removed bool) {
	mgr.muxConns.Lock()
	defer mgr.muxConns.Unlock()

	p, ok := mgr.Conns[addr]
	if !ok {
		return false
	}

	if p.Inbound {
		mgr.NumInbound--
	} else {
		mgr.NumOutbound--
	}

	delete(mgr.Conns, addr)
	return true
}

// AddConn attempts to add a connection with p to mgr's connection pool.
func (mgr *ConnManager) AddConn(p *Peer) (added bool) {
	mgr.muxConns.Lock()
	defer mgr.muxConns.Unlock()
	addr := p.Conn.RemoteAddr().String()

	if p.Inbound {
		if mgr.NumInbound < MaxInboundConns {
			mgr.Conns[addr] = p
			mgr.NumInbound++
			added = true
		} else {
			log.Printf("unable to establish inbound conn with %s", addr)
		}
	} else {
		if mgr.NumOutbound < MaxOutboundConns {
			mgr.Conns[addr] = p
			mgr.NumOutbound++
			added = true
		} else {
			log.Printf("unable to establish outbound conn with %s", addr)
		}
	}

	return added
}

// ClearConns removes all peer connections from mgr's connection pool.
func (mgr *ConnManager) ClearConns() {
	mgr.muxConns.Lock()
	defer mgr.muxConns.Unlock()
	mgr.Conns = make(map[string]*Peer)
}

// SetServices sets the service flags for the peer at addr to svc.
func (mgr *ConnManager) SetServices(addr string, svc protocol.ServiceFlag) {
	p := mgr.GetConn(addr)
	if p != nil {
		p.Services = svc
	}
}
