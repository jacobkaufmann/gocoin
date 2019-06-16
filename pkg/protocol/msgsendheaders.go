package protocol

import "io"

// MsgSendHeaders requests new block headers be announced by the headers
// message as opposed to the inventory message.
type MsgSendHeaders struct{}

// NewMsgSendHeaders returns a new send headers message.
func NewMsgSendHeaders() *MsgSendHeaders {
	return &MsgSendHeaders{}
}

// Serialize serializes msg and writes to w.
func (msg *MsgSendHeaders) Serialize(w io.Writer, pver uint32) error {
	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgSendHeaders) Deserialize(r io.Reader, pver uint32) error {
	return nil
}

// Command returns the message type of the send headers message.
func (msg *MsgSendHeaders) Command() MsgType {
	return MsgTypeSendHeaders
}

// MaxPayloadSize returns the maximum size in bytes of the send headers
// message.
func (msg *MsgSendHeaders) MaxPayloadSize(pver uint32) uint32 {
	return EmptyPayloadSize
}
