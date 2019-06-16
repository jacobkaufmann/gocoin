package protocol

import "io"

// A MsgPong is a response to a ping message, indicating the ponging node is
// still alive. In modern protocol versions, a pong response is generated using
//  a nonce included in the corresponding ping.
type MsgPong struct {
	Nonce uint64
}

// NewMsgPong returns a new pong message containing a specified nonce.
func NewMsgPong(nonce uint64) *MsgPong {
	return &MsgPong{
		Nonce: nonce,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgPong) Serialize(w io.Writer, pver uint32) error {
	return writeElement(w, msg.Nonce)
}

// Deserialize deserializes data from r into msg.
func (msg *MsgPong) Deserialize(r io.Reader, pver uint32) error {
	return readElement(r, &msg.Nonce)
}

// Command returns the message type of the pong message.
func (msg *MsgPong) Command() MsgType {
	return MsgTypePong
}

// MaxPayloadSize returns the maximum size in bytes of the pong message.
func (msg *MsgPong) MaxPayloadSize(pver uint32) uint32 {
	return 8
}
