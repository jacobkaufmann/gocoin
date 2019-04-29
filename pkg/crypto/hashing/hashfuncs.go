package hashing

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

// SHA256B performs the SHA-256 hashing algorithm and returns the corresponding
// bytes.
func SHA256B(b []byte) []byte {
	hash := sha256.Sum256(b)
	return hash[:]
}

// SHA256H performs the SHA-256 hashing algorithm and returns the corresponding
// Hash.
func SHA256H(b []byte) Hash {
	return Hash(sha256.Sum256(b))
}

// DoubleSHA256B performs the SHA-256 hashing algorithm twice sequentially and
// returns the corresponding bytes.
func DoubleSHA256B(b []byte) []byte {
	hash := SHA256B(SHA256B(b))
	return hash[:]
}

// DoubleSHA256H performs the SHA-256 hashing algorithm twice sequentially and
// returns the corresponding Hash.
func DoubleSHA256H(b []byte) Hash {
	return SHA256H(SHA256B(b))
}

// Hash160 performs the Hash160 hashing algorithm and returns the corresonding
// bytes.
func Hash160(b []byte) []byte {
	hash := SHA256B(b)
	hash = ripemd160.New().Sum(hash)
	return hash
}
