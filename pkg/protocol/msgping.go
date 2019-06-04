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

// Serialize converts the ping message into bytes as specified by the protocol
// version pver and writes those bytes to w.
func (msg *MsgPing) Serialize(w io.Writer, pver uint32) error {
	buf := make([]byte, 8)
	littleEndian.PutUint64(buf, msg.Nonce)
	_, err := w.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// Deserialize reads from r and converts those bytes into a ping message as
// specified by the protocol version pver.
func (msg *MsgPing) Deserialize(r io.Reader, pver uint32) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	msg.Nonce = littleEndian.Uint64(buf)
	return nil
}

// Command returns the message type of the ping message.
func (msg *MsgPing) Command() MsgType {
	return MsgTypePing
}

// MaxPayloadSize returns the maximum size in bytes of the ping message.
func (msg *MsgPing) MaxPayloadSize(pver uint32) uint32 {
	return 8
}
