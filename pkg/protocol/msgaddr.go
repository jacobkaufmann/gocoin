package protocol

// MsgAddr is sent to provide information on known network nodes.
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

// MaxPayloadLength returns the maximum length in bytes of the address message.
func (msg *MsgAddr) MaxPayloadLength(pver uint32) uint32 {
	return msg.Count().Size() + NetAddressSize*MaxAddrToSend
}
