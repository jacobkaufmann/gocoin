package p2p

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
)

type inventory struct {
	typeID typeIdentifier
	hash   *hashing.Hash
}

type typeIdentifier uint32

const (
	err      typeIdentifier = 0
	msgTx    typeIdentifier = 1
	msgBlock typeIdentifier = 2
)
