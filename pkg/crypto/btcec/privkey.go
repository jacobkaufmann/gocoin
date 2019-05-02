package btcec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"github.com/jacobkaufmann/gocoin/pkg/util/encoding/base58"
)

// A PrivateKey wraps an ecdsa.PrivateKey and represents a Bitcoin private key.
type PrivateKey ecdsa.PrivateKey

// NewPrivateKey generates a new private key for the elliptic curve.
func NewPrivateKey(curve elliptic.Curve) (*PrivateKey, error) {
	k, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	return (*PrivateKey)(k), nil
}

// PrivKeyFromBytes returns a private and public key for `curve' based on the
// private key passed as an argument as a byte slice.
func PrivKeyFromBytes(curve elliptic.Curve, pk []byte) (*PrivateKey,
	*PublicKey) {
	x, y := curve.ScalarBaseMult(pk)

	priv := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(pk),
	}

	return (*PrivateKey)(priv), (*PublicKey)(&priv.PublicKey)
}

// PubKey returns the PublicKey corresponding to this private key.
func (p *PrivateKey) PubKey() *PublicKey {
	return (*PublicKey)(&p.PublicKey)
}

// Sign generates a signature for the provided hash using the private key.
func (p *PrivateKey) Sign(hash []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, p.ToECDSA(), hash)
	if err != nil {
		return nil, err
	}
	return &Signature{r, s}, nil
}

// ToECDSA returns the private key as a *ecdsa.PrivateKey.
func (p *PrivateKey) ToECDSA() *ecdsa.PrivateKey {
	return (*ecdsa.PrivateKey)(p)
}

// PrivKeyBytesLen defines the length in bytes of a serialized private key.
const PrivKeyBytesLen = 32

// Serialize returns the private key number d as a big-endian binary-encoded
// number, padded to a length of 32 bytes.
func (p *PrivateKey) Serialize() []byte {
	b := make([]byte, 0, PrivKeyBytesLen)
	return paddedAppend(PrivKeyBytesLen, b, p.ToECDSA().D.Bytes())
}

// WIF returns the WIF encoding of the PrivateKey.
func (p *PrivateKey) WIF() string {
	b := p.Serialize()
	return base58.EncodeCheck(b, byte(base58.PrivateKeyWIF))
}
