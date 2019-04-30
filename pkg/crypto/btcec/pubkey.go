package btcec

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"

	"github.com/jacobkaufmann/gocoin/pkg/util/encoding/base58"
)

// A PublicKey represents a Bitcoin public key.
type PublicKey struct {
	*ecdsa.PublicKey
	compressed bool
}

// NewPublicKey returns the public key corresponding to a PrivateKey.
func NewPublicKey(privKey *PrivateKey) *PublicKey {
	pubKey := &PublicKey{
		&privKey.PublicKey,
		privKey.compressed,
	}
	return pubKey
}

// Compressed returns whether or not the PublicKey is compressed.
func (k *PublicKey) Compressed() bool {
	return k.compressed
}

// Address returns the Bitcoin address for a public key.
func (k *PublicKey) Address() string {
	b := []byte(k.Hex())
	h := hashing.Hash160(b)
	return base58.EncodeCheck(h, byte(base58.Address))
}

// Hex returns the hex encoding of the PublicKey.
func (k *PublicKey) Hex() string {
	bigZero := big.NewInt(0)
	xb := k.X.Bytes()
	var h string
	if k.compressed {
		h = hex.EncodeToString(xb)
		if k.Y.Cmp(bigZero) == -1 {
			h = "03" + h
		} else {
			h = "02" + h
		}
	} else {
		b := append(xb, k.Y.Bytes()...)
		h = "04" + string(b[:])
	}
	return h
}
