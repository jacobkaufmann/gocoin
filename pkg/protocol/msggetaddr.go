package protocol

// MsgGetAddr requests information about known active peers to help with
// finding potential nodes in the network.
type MsgGetAddr struct{}

// NewMsgGetAddr returns a new getaddr message.
func NewMsgGetAddr() *MsgGetAddr {
	return &MsgGetAddr{}
}

// Command returns the message type of the getaddr message.
func (msg *MsgGetAddr) Command() MsgType {
	return MsgTypeGetAddr
}

// MaxPayloadSize returns the maximum size in bytes of the getaddr message.
func (msg *MsgGetAddr) MaxPayloadSize(pver uint32) uint32 {
	return EmptyPayloadSize
}
