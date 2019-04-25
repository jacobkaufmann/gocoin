package base58

import (
	"math/big"
)

// Fields for big integer math.
var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Encode encodes a byte slice into Base58.
func Encode(b []byte) string {
	x := new(big.Int)
	x.SetBytes(b)

	// Allocate space for big-endian base58 representation.
	// log(256) / log(58) rounded up.
	encoded := make([]byte, 0, len(b)*136/100)

	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		encoded = append(encoded, alphabet[mod.Int64()])
	}

	// Account for leading zero bytes.
	for _, i := range b {
		if i != 0 {
			break
		}
		encoded = append(encoded, alphabetIdx0)
	}

	// Reverse the string.
	l := len(encoded)
	for i := 0; i < l/2; i++ {
		encoded[i], encoded[l-1-i] = encoded[l-1-i], encoded[i]
	}

	return string(encoded)
}
