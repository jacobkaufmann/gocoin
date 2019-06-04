package protocol

// A MsgGetData requests one or more data objects. The objects are requested
// by an inventory, which the requesting node typically received previously
// by way of an inventory message.
type MsgGetData struct {
	Inventory []*InvVect
}

// NewMsgGetData returns a new getdata message containing a set of inventory
// vectors.
func NewMsgGetData(inv []*InvVect) *MsgGetData {
	return &MsgGetData{
		Inventory: inv,
	}
}

// Count returns the number of inventory entries in the getdata message.
func (msg *MsgGetData) Count() CompactSize {
	return CompactSize(len(msg.Inventory))
}

// Command returns the message type of the getdata message.
func (msg *MsgGetData) Command() MsgType {
	return MsgTypePing
}

// MaxPayloadLength returns the maximum length in bytes of the getdata message.
func (msg *MsgGetData) MaxPayloadLength(pver uint32) uint32 {
	return MaxInvSize*InvVectSize + 9
}
