package protocol

import "io"

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

// Serialize serializes msg and writes to w.
func (msg *MsgNotFound) Serialize(w io.Writer, pver uint32) error {
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
func (msg *MsgNotFound) Deserialize(r io.Reader, pver uint32) error {
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

// Count returns the number of inventory entries in the notfound message.
func (msg *MsgNotFound) Count() uint64 {
	return len(msg.Inventory)
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
