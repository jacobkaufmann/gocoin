package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
)

// A PrivateKey represents the private key in a key pair.
type PrivateKey struct {
	*ecdsa.PrivateKey
	Compressed bool
}

// NewPrivateKey generates a new private key for the elliptic curve.
func NewPrivateKey(compressed bool) *PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	privKey := &PrivateKey{privateKey, compressed}
	return privKey
}

// Hex returns the hex encoding of the PrivateKey.
func (k *PrivateKey) Hex() string {
	h := hex.EncodeToString(k.D.Bytes())
	if k.Compressed {
		h += "01"
	}
	return h
}
