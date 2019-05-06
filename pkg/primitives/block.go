package primitives

import (
	"encoding/binary"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
)

// HeaderSize of array used to store serialized block header.
const HeaderSize = 80

// A Block represents a block in the Bitcoin block chain
type Block struct {
	Size            uint32
	Header          BlockHeader
	NumTransactions uint32
	Transactions    []*Transaction
}

// A BlockHeader contains metadata for a block and is hashed as part
// of Bitcoin's proof of work algorithm.
type BlockHeader struct {
	// The block version number.
	Version uint32

	// The hash of the previous block's header.
	PrevBlockHash hashing.Hash

	// The hash of the merkle root.
	MerkleRootHash hashing.Hash

	// Unix epoch time when the miner started hashing the header
	// (according to the miner).
	Time uint32

	// An encoded version of the 256-bit target threshold this block's header
	// hash must be less than or equal to.
	NumBits uint32

	// An arbitrary number miners change to modify the header hash in order to
	// produce a hash less than or equal to the target threshold.
	// If all 32-bit values are tested, the time can be updated or the coinbase
	// transaction can be changed and the merkle root updated.
	Nonce uint32
}

// Serialize returns the contents of the block header encoded into a single byte
// array.
func (h *BlockHeader) Serialize() [HeaderSize]byte {
	var b [HeaderSize]byte
	binary.LittleEndian.PutUint32(b[:4], h.Version)
	copy(b[4:36], h.PrevBlockHash.Bytes())
	copy(b[36:68], h.MerkleRootHash.Bytes())
	binary.LittleEndian.PutUint32(b[68:72], h.Time)
	binary.LittleEndian.PutUint32(b[72:76], h.NumBits)
	binary.LittleEndian.PutUint32(b[76:], h.Nonce)
	return b
}

// Deserialize converts a serialized byte array into a block header object and
// returns it.
func Deserialize(b [HeaderSize]byte) (*BlockHeader, error) {
	h := &BlockHeader{}
	h.Version = binary.LittleEndian.Uint32(b[:4])
	err := h.PrevBlockHash.SetBytes(b[4:36])
	if err != nil {
		return nil, err
	}
	err = h.MerkleRootHash.SetBytes(b[36:68])
	if err != nil {
		return nil, err
	}
	h.Time = binary.LittleEndian.Uint32(b[68:72])
	h.NumBits = binary.LittleEndian.Uint32(b[72:76])
	h.Nonce = binary.LittleEndian.Uint32(b[76:])
	return h, nil
}
