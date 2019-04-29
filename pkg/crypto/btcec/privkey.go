package btcec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"

	"github.com/jacobkaufmann/gocoin/pkg/util/encoding/base58"
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

// Compressed returns whether or not the PrivateKey is compressed.
func (k *PrivateKey) Compressed() bool {
	return k.compressed
}

// Hex returns the hex encoding of the PrivateKey.
func (k *PrivateKey) Hex() string {
	h := hex.EncodeToString(k.D.Bytes())
	if k.compressed {
		h += "01"
	}
	return h
}

// WIF returns the WIF encoding of the PrivateKey.
func (k *PrivateKey) WIF() string {
	b := []byte(k.Hex())
	return base58.EncodeCheck(b, byte(base58.PrivateKeyWIF))
}
