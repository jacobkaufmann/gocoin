package blockchain

import (
	"math"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
	"github.com/jacobkaufmann/gocoin/pkg/primitives"
)

// MerkleTree represents a Merkle Tree.  The Merkle Tree may be partial
// or whole.
type MerkleTree struct {
	Nodes           []*hashing.Hash
	NumTransactions uint32
}

// HashMerkleNodes takes two hashes, treated as left and right tree nodes, and
// returns the hash of their concatenation.
func HashMerkleNodes(left *hashing.Hash, right *hashing.Hash) *hashing.Hash {
	// Concatenate the left and right nodes.
	var h [64]byte
	copy(h[:len(left)], left[:])
	copy(h[len(left):], right[:])

	newHash := hashing.DoubleSHA256H(h[:])
	return &newHash
}

// nextPowerOfTwo returns the next highest power of two from a given number if
// it is not already a power of two.
func nextPowerOfTwo(n int) int {
	// Return the number if it's already a power of 2.
	if n&(n-1) == 0 {
		return n
	}

	// Figure out and return the next power of two.
	exponent := uint(math.Log2(float64(n))) + 1
	return 1 << exponent // 2^exponent
}

// BuildMerkleTree takes a slice of transaction IDs and returns the MerkleTree built
// from the bottom up.
func BuildMerkleTree(transactions []primitives.Transaction) *MerkleTree {
	nextPoT := nextPowerOfTwo(len(transactions))
	arraySize := nextPoT*2 - 1
	nodes := make([]*hashing.Hash, arraySize)

	// Create the intial bottom layer of hashes and insert into array.
	for i, tx := range transactions {
		// Coinbase transaction
		if i == 0 {
			var coinbaseHash hashing.Hash
			nodes[i] = &coinbaseHash
		} else {
			nodes[i] = &tx.Metadata.TxID
		}
	}

	// Start the array offset after the last transaction and adjusted to the
	// next power of two.
	offset := nextPoT
	for i := 0; i < arraySize-1; i += 2 {
		switch {
		// When there is no left child node, the parent is nil too.
		case nodes[i] == nil:
			nodes[offset] = nil

		// When there is no right child, the parent is generated by
		// hashing the concatenation of the left child with itself.
		case nodes[i+1] == nil:
			newHash := HashMerkleNodes(nodes[i], nodes[i])
			nodes[offset] = newHash

		// The normal case sets the parent node to the double sha256
		// of the concatentation of the left and right children.
		default:
			newHash := HashMerkleNodes(nodes[i], nodes[i+1])
			nodes[offset] = newHash
		}
		offset++
	}

	t := &MerkleTree{}
	t.Nodes = nodes
	t.NumTransactions = uint32(len(transactions))
	return t
}
