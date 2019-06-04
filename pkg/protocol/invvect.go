package protocol

// InvVect defines the structure of an inventory vector. Inventory vectors
// are used for notifying other nodes about objects they have or data which is
// being requested.
type InvVect struct {
	TypeID InvType
	Hash   *[HashSize]byte
}

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
