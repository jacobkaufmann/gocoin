package btcec

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
	"github.com/jacobkaufmann/gocoin/pkg/util/encoding/base58"
)

// These constants define the lengths of serialized public keys.
const (
	PubKeyBytesLenCompressed   = 33
	PubKeyBytesLenUncompressed = 65
)

// These constants define the prefixes for serialized public keys.
const (
	pubkeyCompressedEvenY byte = 0x2
	pubKeyCompressedOddY  byte = 0x3
	pubkeyUncompressed    byte = 0x4
)

// isOdd returns whether a big.Int is odd.  It is a helper for determining
// the appropriate prefix for a serialized public key.
func isOdd(a *big.Int) bool {
	return a.Bit(0) == 1
}

// IsCompressedPubKey returns true the the passed serialized public key has
// been encoded in compressed format, and false otherwise.
func IsCompressedPubKey(pubKey []byte) bool {
	// The public key is only compressed if it is the correct length and
	// the format (first byte) is one of the compressed pubkey values.
	return len(pubKey) == PubKeyBytesLenCompressed &&
		(pubKey[0]&^byte(0x1) == pubkeyCompressedEvenY ||
			pubKey[0]&^byte(0x1) == pubKeyCompressedOddY)
}

// AddressForPubKey returns the Bitcoin address for a public key.
func AddressForPubKey(pubKey []byte) string {
	h := hashing.Hash160(pubKey)
	return base58.EncodeCheck(h, byte(base58.Address))
}

// A PublicKey wraps an ecdsa.PublicKey represents a Bitcoin public key.
type PublicKey ecdsa.PublicKey

// ToECDSA returns the public key as a *ecdsa.PublicKey.
func (p *PublicKey) ToECDSA() *ecdsa.PublicKey {
	return (*ecdsa.PublicKey)(p)
}

// IsEqual compares this public key instance to the one passed, returning true
// if both public keys are equivalent. A public key is equivalent to another,
// if they both have the same X and Y coordinate.
func (p *PublicKey) IsEqual(otherPubKey *PublicKey) bool {
	return p.X.Cmp(otherPubKey.X) == 0 &&
		p.Y.Cmp(otherPubKey.Y) == 0
}

// SerializeUncompressed serializes a public key in a 65-byte uncompressed
// format.
func (p *PublicKey) SerializeUncompressed() []byte {
	b := make([]byte, 0, PubKeyBytesLenUncompressed)
	b = append(b, pubkeyUncompressed)
	b = paddedAppend(32, b, p.X.Bytes())
	return paddedAppend(32, b, p.Y.Bytes())
}

// SerializeCompressed serializes a public key in a 33-byte compressed format.
func (p *PublicKey) SerializeCompressed() []byte {
	var prefix byte
	if isOdd(p.Y) {
		prefix = pubKeyCompressedOddY
	} else {
		prefix = pubkeyCompressedEvenY
	}

	b := make([]byte, 0, PubKeyBytesLenCompressed)
	b = append(b, prefix)
	return paddedAppend(32, b, p.X.Bytes())
}

// paddedAppend appends the src byte slice to dst, returning the new slice.
// If the length of the source is smaller than the passed size, leading zero
// bytes are appended to the dst slice before appending src.
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
