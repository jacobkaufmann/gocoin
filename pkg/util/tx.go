package util

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"

	"github.com/jacobkaufmann/gocoin/pkg/script"
)

// A Tx represents a bitcoin transaction.
type Tx struct {
	version    int32
	txInCount  uint32
	txOutCount uint32
	lockTime   uint32
	size       uint32
	inputs     []*TxIn
	outputs    []*TxOut
}

// A TxID is the unique identifier for a transaction.
type TxID hashing.Hash

// A TxIn is an input to a transaction.
type TxIn struct {
	// Previous transaction output being spent
	prevOutput TxOutPoint

	// Number of bytes in ScriptSig
	scriptBytes uint16

	// A Script which satisfies the conditions specified by the ScriptPubKey in
	// PrevOutput
	scriptSig script.Script

	// Sequence number
	sequence uint32
}

// A TxOut is an output of a transaction.
type TxOut struct {
	// The number of satoshis to spend
	// One satoshi is 10^-8 of a bitcoin
	value uint64

	// Number of bytes in PubKeyScript
	scriptPubKeyBytes uint16

	// A Script which defines conditions which must be satisfied to spend this
	// transaction output
	scriptPubKey script.Script
}

// A TxOutPoint contains information to refer to a specific
// transaction output.
type TxOutPoint struct {
	// The identifier for the transaction holding the transaction output
	hash TxID

	// The index of the transaction output in the transaction referenced by Hash
	index uint32
}
