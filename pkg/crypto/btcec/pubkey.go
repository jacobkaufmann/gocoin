package btcec

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
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
