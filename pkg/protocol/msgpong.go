package protocol

import "io"

// A MsgPong is sent in response to a ping message.  In modern protocol
// versions, a pong response is generated using a nonce included in the
// corresponding ping.
type MsgPong struct {
	Nonce uint64
}

// NewMsgPong returns a new pong message containing a specified nonce.
func NewMsgPong(nonce uint64) *MsgPong {
	return &MsgPong{
		Nonce: nonce,
	}
}

// Serialize converts the pong message into bytes as specified by the protocol
// version pver and writes those bytes to w.
func (msg *MsgPong) Serialize(w io.Writer, pver uint32) error {
	buf := make([]byte, 8)
	littleEndian.PutUint64(buf, msg.Nonce)
	_, err := w.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// Deserialize reads from r and converts those bytes into a pong message as
// specified by the protocol version pver.
func (msg *MsgPong) Deserialize(r io.Reader, pver uint32) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	msg.Nonce = littleEndian.Uint64(buf)
	return nil
}

// Command returns the message type of the pong message.
func (msg *MsgPong) Command() MsgType {
	return MsgTypePong
}

// MaxPayloadLength returns the maximum length in bytes of the pong message.
func (msg *MsgPong) MaxPayloadLength(pver uint32) uint32 {
	return 8
}
