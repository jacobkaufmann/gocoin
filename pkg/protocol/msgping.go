package protocol

// A MsgPing is sent to confirm the TCP/IP connection is still valid.
type MsgPing struct {
	nonce uint64
}

func (msg *MsgPing) Command() MsgType {
	return MsgTypePing
}

func (msg *MsgPing) MaxPayloadLength(pver uint32) uint32 {
	return 8
}
