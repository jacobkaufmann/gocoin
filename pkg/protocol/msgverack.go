package protocol

// MsgVerAck acknowledges a previously received version message.
type MsgVerAck struct{}

// Command returns the message type of the verack message.
func (msg *MsgVerAck) Command() MsgType {
	return MsgTypeVerAck
}

// MaxPayloadSize returns the maximum size in bytes of the verack message.
func (msg *MsgVerAck) MaxPayloadSize(pver uint32) uint32 {
	return EmptyPayloadSize
}
