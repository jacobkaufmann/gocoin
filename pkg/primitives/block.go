package primitives

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
)

// A Block represents a block in the Bitcoin block chain
type Block struct {
	Size         uint32
	Header       BlockHeader
	Transactions []*Transaction
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
