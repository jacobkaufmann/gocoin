package protocol

import (
	"net"
	"time"
)

// NetAddress defines the structure used to represent network addresses.
// Network addresses are not prefixed with timestamp in the version message.
type NetAddress struct {
	Timestamp time.Time
	Services  ServiceFlag
	IP        net.IP
	Port      uint16
}

// NetAddressSize is the size in bytes of a network address.
const NetAddressSize = 30
