package protocol

import "io"

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

// Serialize serializes msg and writes to w.
func (msg *MsgGetData) Serialize(w io.Writer, pver uint32) error {
	err := writeCompactSize(w, pver, msg.InvCount())
	if err != nil {
		return err
	}

	for _, inv := range msg.Inventory {
		err := inv.Serialize(w, pver)
		if err != nil {
			return err
		}
	}

	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgGetData) Deserialize(r io.Reader, pver uint32) error {
	var n uint64
	err := readCompactSize(r, pver, &n)
	if err != nil {
		return err
	}

	for i := 0; i < int(n); i++ {
		inv := &InvVect{}
		err = inv.Deserialize(r, pver)
		if err != nil {
			return err
		}
		msg.Inventory = append(msg.Inventory, inv)
	}

	return nil
}

// InvCount returns the number of inventory entries in the getdata message.
func (msg *MsgGetData) InvCount() uint64 {
	return uint64(len(msg.Inventory))
}

// Command returns the message type of the getdata message.
func (msg *MsgGetData) Command() MsgType {
	return MsgTypePing
}

// MaxPayloadLength returns the maximum length in bytes of the getdata message.
func (msg *MsgGetData) MaxPayloadLength(pver uint32) uint32 {
	return MaxInvSize*InvVectSize + 9
}
