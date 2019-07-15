package util

import (
	"math"
	"time"

	"github.com/jacobkaufmann/gocoin/pkg/protocol"
)

// A Block represents a block in the bitcoin blockchain. A Block is a higher
// level representation of a block message with additional facilities.
type Block struct {
	*protocol.BlockHeader
	Txns []*Tx
}

// NewBlock returns an initialized block for a set of transactions.
func NewBlock(version uint32, prevBlockHash *[protocol.HashSize]byte,
	timestamp time.Time, numBits uint32, txns []*Tx) *Block {
	// TODO: resolve use of hashing.Hash and [HashSize]byte.
	blk := &Block{}
	blk.Version = version
	blk.PrevBlockHash = prevBlockHash
	blk.Timestamp = timestamp
	blk.NumBits = numBits
	blk.Txns = txns
	return blk
}

// NewBlockFromMsg returns an initialized block from a block message.
func NewBlockFromMsg(msg *protocol.MsgBlock) *Block {
	blk := &Block{}
	blk.BlockHeader = msg.BlockHeader
	for _, tx := range msg.Txns {
		blk.Txns = append(blk.Txns, NewTxFromMsg(tx))
	}
	return blk
}

// maxNonce is a convenience variable representing the maximum possible nonce
// value.
const maxNonce = math.MaxUint32

// IncrementNonce increments the block's nonce. If the nonce is at its maximum
// value, the nonce goes to zero and the reset flag is set to true to notify
// the miner that all nonce values have been exhausted.
func (blk *Block) IncrementNonce() (reset bool) {
	if blk.Nonce == maxNonce {
		reset = true
	}
	blk.Nonce++
	return
}
