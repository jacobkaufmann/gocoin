package protocol

import "io"

// MsgMempool requests information about transactions a node has verified but
// which are yet to be confirmed.
type MsgMempool struct{}

// NewMsgMempool returns a new mempool message.
func NewMsgMempool() *MsgMempool {
	return &MsgMempool{}
}

// Serialize serializes msg and writes to w.
func (msg *MsgMempool) Serialize(w io.Writer, pver uint32) error {
	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgMempool) Deserialize(r io.Reader, pver uint32) error {
	return nil
}

// Command returns the message type of the mempool message.
func (msg *MsgMempool) Command() MsgType {
	return MsgTypeMempool
}

// MaxPayloadSize returns the maximum size in bytes of the mempool message.
func (msg *MsgMempool) MaxPayloadSize(pver uint32) uint32 {
	return EmptyPayloadSize
}
