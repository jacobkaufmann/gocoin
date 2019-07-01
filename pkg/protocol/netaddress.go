package protocol

import (
	"io"
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

// writeNetAddress encodes addr and writes the value to w.
func writeNetAddress(w io.Writer, pver uint32, addr NetAddress,
	includeTime bool) error {
	var err error

	if includeTime {
		timestamp := uint32Time(addr.Timestamp)
		err = writeElement(w, timestamp)
		if err != nil {
			return err
		}
	}

	return writeElements(w, addr.Services, addr.IP.To16(), addr.Port)
}

// readNetAddress reads from r and decodes the value into addr.
func readNetAddress(r io.Reader, pver uint32, addr *NetAddress,
	includeTime bool) error {
	var err error

	if includeTime {
		var timestamp uint32Time
		err = readElements(r, &timestamp)
		if err != nil {
			return err
		}
		addr.Timestamp = time.Time(timestamp)
	}

	var ip [net.IPv6len]byte
	err = readElements(r, &addr.Services, &ip, &addr.Port)
	if err != nil {
		return err
	}

	addr.IP = ip[:]
	return nil
}
