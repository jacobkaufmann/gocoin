package protocol

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

// Count returns the number of addresses in the address message.
func (msg *MsgAddr) Count() CompactSize {
	return CompactSize(len(msg.Addrs))
}

// Command returns the message type of the address message.
func (msg *MsgAddr) Command() MsgType {
	return MsgTypeAddr
}

// MaxPayloadSize returns the maximum size in bytes of the address message.
func (msg *MsgAddr) MaxPayloadSize(pver uint32) uint32 {
	return msg.Count().Size() + NetAddressSize*MaxAddrToSend
}
