package util

import (
	"math"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
)

// HeaderSize of array used to store serialized block header.
const HeaderSize = 80

// A Block represents a block in the bitcoin block chain.
type Block struct {
	// The block version number.
	version uint32

	// The hash of the previous block's header.
	prevBlockHash *hashing.Hash

	// The hash of the merkle root.
	merkleRootHash *hashing.Hash

	// Unix epoch time when the miner started hashing the header
	// (according to the miner).
	time uint32

	// An encoded version of the 256-bit target threshold this block's header
	// hash must be less than or equal to.
	numBits uint32

	// An arbitrary number miners change to modify the header hash in order to
	// produce a hash less than or equal to the target threshold.
	// If all 32-bit values are tested, the time can be updated or the coinbase
	// transaction can be changed and the merkle root updated.
	nonce uint32

	// Number of transaction entries.
	txCount CompactSize

	// The transactions included in the block.
	txns []*Tx
}

// Header assembles and returns the block header corresponding to a block.
func (blck *Block) Header() *BlockHeader {
	return &BlockHeader{
		version:        &blck.version,
		prevBlockHash:  blck.prevBlockHash,
		merkleRootHash: blck.merkleRootHash,
		time:           &blck.time,
		numBits:        &blck.numBits,
		nonce:          &blck.nonce,
		txCount:        &blck.txCount,
	}
}

// maxNonce is a convenience variable representing the maximum possible nonce
// value.
const maxNonce = math.MaxUint32

// IncrementNonce increments the block's nonce.  If the nonce is at its maximum
// value, the nonce goes to zero and the reset flag is set to true to notify
// the miner that all nonce values have been exhausted.
//
// Incrementing the nonce is fundamental to the mining process.
func (blck *Block) IncrementNonce() (reset bool) {
	if blck.nonce == maxNonce {
		reset = true
	}
	blck.nonce++
	return
}

// A BlockHeader contains metadata for a block and is hashed as part
// of bitcoin's proof of work algorithm.
//
// A BlockHeader functions as a view of a block.
type BlockHeader struct {
	version        *uint32
	prevBlockHash  *hashing.Hash
	merkleRootHash *hashing.Hash
	time           *uint32
	numBits        *uint32
	nonce          *uint32
	txCount        *CompactSize
}
