package mining

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
)

const (
	coinbaseOutputIndex = 0xFFFFFFFF
	coinbaseSequence    = 0xFFFFFFFF
)

var (
	coinbasePrevTxHash = hashing.Hash([hashing.HashSize]byte{0})
)
