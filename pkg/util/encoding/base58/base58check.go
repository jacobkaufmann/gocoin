package base58

import (
	"errors"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
)

// VersionPrefix represents a Base58Check version prefix.
type VersionPrefix byte

const (
	// Address refers to a Bitcoin address.
	Address VersionPrefix = 0
	// P2SH refers to a Pay-to-Script-Hash address.
	P2SH VersionPrefix = 5
	// PrivateKeyWIF refers to a Private key WIF (Wallet Import Format).
	PrivateKeyWIF VersionPrefix = 128
)

var (
	// ErrCheckSum  indicates that the checksum of a check-encoded string
	// does not verify against the checksum.
	ErrCheckSum = errors.New("checksum error")

	// ErrInvalidFormat indicates that the check-encoded string has an
	// invalid format.
	ErrInvalidFormat = errors.New("invalid format: version and/or checksum missing")
)

// checksum computes the 32-bit checksum for a byte slice.
func checksum(b []byte) (check [4]byte) {
	h := hashing.DoubleSHA256B(b)
	copy(check[:], h[:4])
	return
}

// EncodeCheck encodes a byte slice into Base58Check.
func EncodeCheck(payload []byte, version byte) string {
	b := make([]byte, 0, 1+len(payload)+4)
	b = append(b, version)
	b = append(b, payload[:]...)
	check := checksum(b)
	b = append(b, check[:]...)
	return Encode(b)
}

// DecodeCheck decodes a Base58Check string into a destination
// byte slice.
func DecodeCheck(src string) (dst []byte, version byte, err error) {
	decoded, err := Decode(src)
	if err != nil {
		return nil, 0, err
	}
	l := len(decoded)
	if l < 5 {
		return nil, 0, ErrInvalidFormat
	}

	// Extract checksum and verify.
	var check [4]byte
	copy(check[:], decoded[l-4:])
	if checksum(decoded[l-4:]) != check {
		return nil, 0, ErrCheckSum
	}

	// Extract version and payload.
	version = decoded[0]
	copy(dst, decoded[1:l-4])
	return
}
