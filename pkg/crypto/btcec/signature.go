package btcec

import (
	"crypto/ecdsa"
	"math/big"
)

// A Signature represents an ECDSA signature.
type Signature struct {
	R *big.Int
	S *big.Int
}

// Verify calls ecdsa.Verify to verify the signature of hash using the public
// key.
func (sig *Signature) Verify(hash []byte, pubKey *PublicKey) bool {
	return ecdsa.Verify(pubKey.ToECDSA(), hash, sig.R, sig.S)
}

// IsEqual compares this Signature instance to the one passed, returning true
// if both Signatures are equivalent. A signature is equivalent to another, if
// they both have the same scalar value for R and S.
func (sig *Signature) IsEqual(otherSig *Signature) bool {
	return sig.R.Cmp(otherSig.R) == 0 &&
		sig.S.Cmp(otherSig.S) == 0
}
