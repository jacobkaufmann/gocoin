package util

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"

	"github.com/jacobkaufmann/gocoin/pkg/script"
)

// A Tx represents a bitcoin transaction.
type Tx struct {
	version    int32
	txInCount  CompactSize
	inputs     []*TxIn
	txOutCount CompactSize
	outputs    []*TxOut
	lockTime   uint32
	size       uint32
}

// NewTx returns an initialized transaction for a set of inputs and outputs.
func NewTx(version int32, inputs []*TxIn, outputs []*TxOut, lockTime uint32) *Tx {
	tx := &Tx{
		version:    version,
		txInCount:  CompactSize(len(inputs)),
		inputs:     inputs,
		txOutCount: CompactSize(len(outputs)),
		outputs:    outputs,
		lockTime:   lockTime,
	}

	// TODO: find a more elegant/safe way to compute this without so many type
	// conversions.
	size := 12 + uint64(len(tx.txInCount.Bytes())) + uint64(len(tx.txOutCount.Bytes()))
	for i := 0; i < int(tx.txInCount.Uint64()); i++ {
		size += uint64(tx.inputs[i].size())
	}
	for i := 0; i < int(tx.txOutCount.Uint64()); i++ {
		size += uint64(tx.outputs[i].size())
	}
	tx.size = uint32(size)

	return tx
}

// Version returns the transaction version number.
func (tx *Tx) Version() int32 {
	return tx.version
}

// TxInCount returns the number of transaction inputs for the transaction.
func (tx *Tx) TxInCount() CompactSize {
	return tx.txInCount
}

// Inputs returns the transaction inputs for the transaction.
func (tx *Tx) Inputs() []*TxIn {
	return tx.inputs
}

// TxOutCount returns the number of transaction outputs for the transaction.
func (tx *Tx) TxOutCount() CompactSize {
	return tx.txOutCount
}

// Outputs returns the transaction outputs for the transaction.
func (tx *Tx) Outputs() []*TxOut {
	return tx.outputs
}

// LockTime returns the lock time of the transaction.
func (tx *Tx) LockTime() uint32 {
	return tx.lockTime
}

// Size returns the size in bytes of the transaction
func (tx *Tx) Size() uint32 {
	return tx.size
}

// A TxID is the unique identifier for a transaction.
type TxID hashing.Hash

// A TxIn is an input to a transaction.
type TxIn struct {
	// Previous transaction output being spent
	prevOutput TxOutPoint

	// Number of bytes in ScriptSig
	scriptBytes CompactSize

	// A Script which satisfies the conditions specified by the ScriptPubKey in
	// PrevOutput
	scriptSig script.Script

	// Sequence number
	sequence uint32
}

func (in *TxIn) size() uint32 {
	// TODO: find a more elegant/safe way to compute this without so many type
	// conversions.
	size := uint64(txOutPointSize)
	size += uint64(len(in.scriptBytes.Bytes())) + in.scriptBytes.Uint64()
	size += txOutPointSize
	return uint32(size)
}

// A TxOut is an output of a transaction.
type TxOut struct {
	// The number of satoshis to spend
	// One satoshi is 10^-8 of a bitcoin
	value uint64

	// Number of bytes in PubKeyScript
	scriptPubKeyBytes CompactSize

	// A Script which defines conditions which must be satisfied to spend this
	// transaction output
	scriptPubKey script.Script
}

func (out *TxOut) size() uint32 {
	// TODO: find a more elegant/safe way to compute this without so many type
	// conversions.
	size := 8 + uint64(len(out.scriptPubKeyBytes.Bytes())) + out.scriptPubKeyBytes.Uint64()
	return uint32(size)
}

// A TxOutPoint contains information to refer to a specific
// transaction output.
type TxOutPoint struct {
	// The identifier for the transaction holding the transaction output
	hash TxID

	// The index of the transaction output in the transaction referenced by Hash
	index uint32
}

const txOutPointSize = hashing.HashSize + 4
