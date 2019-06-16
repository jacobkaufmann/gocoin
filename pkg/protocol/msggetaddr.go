package protocol

import "io"

// MsgGetAddr requests information about known active peers to help with
// finding potential nodes in the network.
type MsgGetAddr struct{}

// NewMsgGetAddr returns a new getaddr message.
func NewMsgGetAddr() *MsgGetAddr {
	return &MsgGetAddr{}
}

// Serialize serializes msg and writes to w.
func (msg *MsgGetAddr) Serialize(w io.Writer, pver uint32) error {
	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgGetAddr) Deserialize(r io.Reader, pver uint32) error {
	return nil
}

// Command returns the message type of the getaddr message.
func (msg *MsgGetAddr) Command() MsgType {
	return MsgTypeGetAddr
}

// MaxPayloadSize returns the maximum size in bytes of the getaddr message.
func (msg *MsgGetAddr) MaxPayloadSize(pver uint32) uint32 {
	return EmptyPayloadSize
}
