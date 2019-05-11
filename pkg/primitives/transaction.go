package primitives

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"

	"github.com/jacobkaufmann/gocoin/pkg/script"
)

// A Tx represents a Bitcoin transaction.
type Tx struct {
	Metadata TxMetadata
	Inputs   []*TxInput
	Outputs  []*TxOutput
}

// TxMetadata holds metadata related to a transaction.
type TxMetadata struct {
	// The hash of the entire transaction
	// Serves as the unique identifier for the transaction
	TxID hashing.Hash

	// Transaction version number
	Version int32

	// Number of inputs in the transaction
	TxInputsCount uint32

	// Number of outputs in the transaction
	TxOutputsCount uint32

	// A time (Unix epoch time) or block number which indicates the earliest time
	// or earliest block when this transaction may be added to the block chain
	LockTime uint32

	// The size of the transaction
	Size uint32
}

// A TxInput is an input to a transaction.
type TxInput struct {
	// Previous TxOutpoint being spent
	PrevOutput TxOutpoint

	// Number of bytes in ScriptSig
	ScriptBytes uint16

	// A Script which satisfies the conditions specified by the ScriptPubKey in
	// PrevOutput
	ScriptSig script.Script

	// Sequence number
	Sequence uint32
}

// A TxOutput is an output of a transaction.
type TxOutput struct {
	// The number of satoshis to spend
	// One satoshi is 10^-8 of a bitcoin
	Value uint64

	// Number of bytes in PubKeyScript
	ScriptPubKeyBytes uint16

	// A Script which defines conditions which must be satisfied to spend this
	// transaction output
	ScriptPubKey script.Script
}

// A TxOutpoint contains information to refer to a specific
// transaction output.
type TxOutpoint struct {
	// The identifier for the transaction holding the transaction output
	Hash hashing.Hash

	// The index of the transaction output in the transaction referenced by Hash
	Index uint32
}
