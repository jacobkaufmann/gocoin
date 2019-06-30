package protocol

import "io"

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

// Serialize serializes msg and writes to w.
func (msg *MsgInv) Serialize(w io.Writer, pver uint32) error {
	err := writeCompactSize(w, pver, msg.Count())
	if err != nil {
		return err
	}

	for _, v := range msg.Inventory {
		err := writeElements(w, v.TypeID, v.Hash)
		if err != nil {
			return err
		}
	}

	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgInv) Deserialize(r io.Reader, pver uint32) error {
	var n uint64
	err := readCompactSize(r, pver, &n)
	if err != nil {
		return err
	}

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

// Count returns the number of inventory entries in the inv message.
func (msg *MsgInv) Count() uint64 {
	return len(msg.Inventory)
}

// Command returns the message type of the inventory message.
func (msg *MsgInv) Command() MsgType {
	return MsgTypeInv
}

// MaxPayloadSize returns the maximum size in bytes of the inventory
// message.
func (msg *MsgInv) MaxPayloadSize(pver uint32) uint32 {
	return MaxInvSize*InvVectSize + 9
}
