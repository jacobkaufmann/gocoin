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
	count := CompactSize(uint64(len(msg.Inventory)))
	err := count.Serialize(w, pver)
	if err != nil {
		return err
	}

	for _, v := range msg.Inventory {
		err := writeElement(w, v.TypeID, v.Hash)
		if err != nil {
			return err
		}
	}

	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgGetData) Deserialize(r io.Reader, pver uint32) error {
	var count CompactSize
	err = count.Deserialize(r, pver)
	if err != nil {
		return err
	}

	n := count.Uint64()
	for i := 0; i < n; i++ {
		invVect := &InvVect{}
		err := readElements(r, &invVect.TypeID, &invVect.Hash)
		if err != nil {
			return err
		}

		msg.Inventory = append(msg.Inventory, invVect)
	}

	return nil
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
