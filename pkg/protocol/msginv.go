package protocol

// A MsgInv transmits one or more inventory vectors of objects known to the
// transmitting peer. It may be sent unsolicited or in response to a getblocks
// message or mempool message.
type MsgInv struct {
	Inventory []*InvVect
}

// NewMsgInv returns a new inventory message containing a set of inventory
// vectors.
func NewMsgInv(inv []*InvVect) *MsgInv {
	return &MsgInv{
		Inventory: inv,
	}
}

// Count returns the number of inventory entries in the inv message.
func (msg *MsgInv) Count() CompactSize {
	return CompactSize(len(msg.Inventory))
}

// Command returns the message type of the inventory message.
func (msg *MsgInv) Command() MsgType {
	return MsgTypeGetAddr
}

// MaxPayloadSize returns the maximum size in bytes of the inventory
// message.
func (msg *MsgInv) MaxPayloadSize(pver uint32) uint32 {
	return MaxInvSize*InvVectSize + 9
}
