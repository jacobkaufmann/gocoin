package protocol

import "io"

// MsgVerAck acknowledges a previously received version message.
type MsgVerAck struct{}

// NewMsgVerAck returns a new verack message.
func NewMsgVerAck() *MsgVerAck {
	return &MsgVerAck{}
}

// Serialize serializes msg and writes to w.
func (msg *MsgVerAck) Serialize(w io.Writer, pver uint32) error {
	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgVerAck) Deserialize(r io.Reader, pver uint32) error {
	return nil
}

// Command returns the message type of the verack message.
func (msg *MsgVerAck) Command() MsgType {
	return MsgTypeVerAck
}

// MaxPayloadSize returns the maximum size in bytes of the verack message.
func (msg *MsgVerAck) MaxPayloadSize(pver uint32) uint32 {
	return EmptyPayloadSize
}
