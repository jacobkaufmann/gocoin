package hashing

import (
	"encoding/hex"
	"fmt"
)

// HashSize of array used to store hashes.
const HashSize = 32

// Hash is used in several of the bitcoin messages and common structures.  It
// typically represents the double sha256 of data.
type Hash [HashSize]byte

// String returns the Hash as the hexadecimal string of the byte-reversed
// hash.
func (hash Hash) String() string {
	for i := 0; i < HashSize/2; i++ {
		hash[i], hash[HashSize-1-i] = hash[HashSize-1-i], hash[i]
	}
	return hex.EncodeToString(hash[:])
}

// Bytes returns the hash as a byte slice.
func (hash *Hash) Bytes() []byte {
	return hash[:]
}

// SetBytes sets the bytes which represent the hash.  An error is returned if
// the number of bytes passed in is not HashSize.
func (hash *Hash) SetBytes(newHash []byte) error {
	nlen := len(newHash)
	if nlen != HashSize {
		return fmt.Errorf("invalid hash length of %v, want %v", nlen,
			HashSize)
	}

	copy(hash[:], newHash)
	return nil
}

// NewHash returns a new Hash from a byte slice.  An error is returned if
// the number of bytes passed in is not HashSize.
func NewHash(newHash []byte) (*Hash, error) {
	var hash Hash
	err := hash.SetBytes(newHash)
	if err != nil {
		return nil, err
	}
	return &hash, err
}
