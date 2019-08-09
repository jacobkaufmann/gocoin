package mining

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
	"github.com/jacobkaufmann/gocoin/pkg/util"
)

const (
	coinbaseOutputIndex = 0xFFFFFFFF
	coinbaseSequence    = 0xFFFFFFFF
)

var (
	coinbasePrevTxHash = hashing.Hash([hashing.HashSize]byte{0})
)

// Miner represents a bitcoin mining node.
type Miner interface {
	Mine(blk *util.Block)
}
