package util

import (
	"bytes"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
	"github.com/jacobkaufmann/gocoin/pkg/protocol"
)

// A Tx represents a bitcoin transaction. A Tx is a higher level representation
// of a transaction message with additional facilities.
type Tx struct {
	*protocol.MsgTx
}

// NewTx returns an initialized transaction for a set of inputs and outputs.
func NewTx(version int32, inputs []*protocol.TxIn, outputs []*protocol.TxOut, lockTime uint32) *Tx {
	msgTx := &protocol.MsgTx{
		Version:  version,
		Inputs:   inputs,
		Outputs:  outputs,
		LockTime: lockTime,
	}
	return &Tx{msgTx}
}

// NewTxFromMessage returns an initialized transaction from a transaction message.
func NewTxFromMessage(msg *protocol.MsgTx) *Tx {
	return &Tx{MsgTx: msg}
}

// Message returns the underlying transaction message of the transaction.
func (tx *Tx) Message() *protocol.MsgTx {
	return tx.MsgTx
}

// TxID returns the transaction id (double-SHA256 hash) of tx.
func (tx *Tx) TxID(pver uint32) (*hashing.Hash, error) {
	var buf bytes.Buffer
	err := tx.MsgTx.Serialize(&buf, pver)
	if err != nil {
		return nil, err
	}

	txID := hashing.DoubleSHA256H(buf.Bytes())
	return &txID, nil
}

// AddInput adds a transaction input to the transaction.
func (tx *Tx) AddInput(in *protocol.TxIn) {
	tx.Inputs = append(tx.Inputs, in)
}

// ClearInputs removes all transaction inputs from the transaction.
func (tx *Tx) ClearInputs() {
	tx.Inputs = nil
}

// AddOutput adds a transaction output to the transaction.
func (tx *Tx) AddOutput(out *protocol.TxOut) {
	tx.Outputs = append(tx.Outputs, out)
}

// ClearOutputs removes all transaction outputs from the transaction.
func (tx *Tx) ClearOutputs() {
	tx.Outputs = nil
}
