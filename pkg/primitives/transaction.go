package primitives

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"

	"github.com/jacobkaufmann/gocoin/pkg/script"
)

// A Transaction represents a bitcoin transaction
type Transaction struct {
	Metadata TransactionMetadata
	Inputs   []*TransactionInput
	Outputs  []*TransactionOutput
}

// TransactionMetadata holds metadata related to the Transaction.
type TransactionMetadata struct {
	// The hash of the entire Transaction
	// Serves as the the unique identifier for the Transaction
	TxID hashing.Hash

	// Transaction version number
	Version uint32

	// Number of inputs in the Transaction
	TxInputsCount uint32

	// Number of outputs in the Transaction
	TxOutputsCount uint32

	// A time (Unix epoch time) or block number which indicates the earliest time
	// or earliest block when this Transaction may be added to the block chain
	LockTime uint32

	// The size of the Transaction
	Size uint32
}

// A TransactionInput is an input to a Transaction.
type TransactionInput struct {
	// Previous TransactionOutpoint being spent
	PrevOutput TransactionOutpoint

	// Number of bytes in ScriptSig
	ScriptBytes uint16

	// A Script which satisfies the conditions specified by the ScriptPubKey in
	// PreviousOutput
	ScriptSig script.Script

	// Sequence number
	Sequence uint32
}

// A TransactionOutput is an output of a Transaction.
type TransactionOutput struct {
	// The number of satoshis to spend
	// One satoshi is 10^-8 of a bitcoin
	Value uint64

	// Number of bytes in PubKeyScript
	ScriptPubKeyBytes uint16

	// A Script which defines conditions which must be satisfied to spend this
	// TransactionOutput
	ScriptPubKey script.Script
}

// A TransactionOutpoint contains information to refer to a specific
// TransactionOutput.
type TransactionOutpoint struct {
	// The identifier for the Transaction holding the TransactionOutput
	Hash hashing.Hash

	// The index of the TransactionOutput in the Transaction referenced by Hash
	Index uint32
}
