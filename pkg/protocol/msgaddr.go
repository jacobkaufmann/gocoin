package protocol

import "io"

// MsgAddr provides information on known network nodes.
type MsgAddr struct {
	Addrs []*NetAddress
}

// NewMsgAddr returns a new address message containing addresses addrs.
func NewMsgAddr(addrs []*NetAddress) *MsgAddr {
	return &MsgAddr{
		Addrs: addrs,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgAddr) Serialize(w io.Writer, pver uint32) error {
	err := writeCompactSize(w, pver, uint64(msg.AddrCount()))
	if err != nil {
		return err
	}

	for _, addr := range msg.Addrs {
		err = writeNetAddress(w, pver, *addr, true)
		if err != nil {
			return err
		}
	}
	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgAddr) Deserialize(r io.Reader, pver uint32) error {
	var n uint64
	err := readCompactSize(r, pver, &n)
	if err != nil {
		return err
	}

	for i := 0; i < int(n); i++ {
		addr := &NetAddress{}
		err = readNetAddress(r, pver, addr, true)
		if err != nil {
			return err
		}
		msg.Addrs = append(msg.Addrs, addr)
	}

	return nil
}

// AddrCount returns the number of addresses in the address message.
func (msg *MsgAddr) AddrCount() uint64 {
	return uint64(len(msg.Addrs))
}

// Command returns the message type of the address message.
func (msg *MsgAddr) Command() MsgType {
	return MsgTypeAddr
}

// MaxPayloadSize returns the maximum size in bytes of the address message.
func (msg *MsgAddr) MaxPayloadSize(pver uint32) uint32 {
	// The conversion to uint32 is safe because of the limit on network
	// addresses.
	return uint32(msg.AddrCount()) + NetAddressSize*MaxAddrToSend
}
