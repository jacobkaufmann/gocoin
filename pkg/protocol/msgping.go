package protocol

import "io"

// A MsgPing confirms the TCP/IP connection to the receiving peer is still
// valid.
type MsgPing struct {
	Nonce uint64
}

// NewMsgPing returns a new ping message containing a specified nonce.
func NewMsgPing(nonce uint64) *MsgPing {
	return &MsgPing{
		Nonce: nonce,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgPing) Serialize(w io.Writer, pver uint32) error {
	return writeElement(w, msg.Nonce)
}

// Deserialize deserializes data from r into msg.
func (msg *MsgPing) Deserialize(r io.Reader, pver uint32) error {
	return readElement(r, &msg.Nonce)
}

// Command returns the message type of the ping message.
func (msg *MsgPing) Command() MsgType {
	return MsgTypePing
}

// MaxPayloadSize returns the maximum size in bytes of the ping message.
func (msg *MsgPing) MaxPayloadSize(pver uint32) uint32 {
	return 8
}
