package protocol

// MsgNotFound is a response to a getdata message which requested an object
// the receiver does not have available for relay.
type MsgNotFound struct {
	Inventory []*InvVect
}

// NewMsgNotFound returns a new notfound message containing a slice of
// inventory vectors.
func NewMsgNotFound(inv []*InvVect) *MsgInv {
	return &MsgInv{
		Inventory: inv,
	}
}

// Count returns the number of inventory entries in the notfound message.
func (msg *MsgNotFound) Count() CompactSize {
	return CompactSize(len(msg.Inventory))
}

// Command returns the message type of the notfound message.
func (msg *MsgNotFound) Command() MsgType {
	return MsgTypeNotFound
}

// MaxPayloadSize returns the maximum size in bytes of the notfound
// message.
func (msg *MsgNotFound) MaxPayloadSize(pver uint32) uint32 {
	return MaxInvSize*InvVectSize + 9
}
