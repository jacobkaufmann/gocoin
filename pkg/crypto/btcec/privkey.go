package btcec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

// A PrivateKey represents a Bitcoin private key.
type PrivateKey struct {
	*ecdsa.PrivateKey
	compressed bool
}

// NewPrivateKey generates a new private key for the elliptic curve.
func NewPrivateKey(compressed bool) *PrivateKey {
	k, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	privKey := &PrivateKey{k, compressed}
	return privKey
}

// Hex returns the hex encoding of the PrivateKey.
func (k *PrivateKey) Hex() string {
	h := hex.EncodeToString(k.D.Bytes())
	if k.compressed {
		h += "01"
	}
	return h
}
