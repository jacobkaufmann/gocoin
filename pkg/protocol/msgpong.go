package protocol

// A MsgPong is sent in response to a ping message.  In modern protocol
// versions, a pong response is generated using a nonce included in the
// corresponding ping.
type MsgPong struct {
	nonce uint64
}

func (msg *MsgPong) Command() MsgType {
	return MsgTypePong
}

func (msg *MsgPong) MaxPayloadLength(pver uint32) uint32 {
	return 8
}
