package protocol

import "io"

// InvType defines the object type linked to an inventory vector.
type InvType uint32

// Bitcoin inventory types.
const (
	InvTypeUndefined InvType = 0
	InvTypeMsgTx     InvType = 1
	InvTypeMsgBlock  InvType = 2
)

// InvVectSize is the size in bytes of an inventory vector.
const InvVectSize = HashSize + 4

// InvVect defines the structure of an inventory vector. Inventory vectors
// are used for notifying other nodes about objects they have or data which is
// being requested.
type InvVect struct {
	TypeID InvType
	Hash   *[HashSize]byte
}

// NewInvVect returns a new inventory vector of a specified type and hash.
func NewInvVect(typeID InvType, hash *[HashSize]byte) *InvVect {
	return &InvVect{
		TypeID: typeID,
		Hash:   hash,
	}
}

// Serialize serializes inv and writes to w.
func (inv *InvVect) Serialize(w io.Writer, pver uint32) error {
	return writeElements(w, inv.TypeID, inv.Hash)
}

// Deserialize deserializes data from r into inv.
func (inv *InvVect) Deserialize(r io.Reader, pver uint32) error {
	return readElements(r, &inv.TypeID, &inv.Hash)
}
