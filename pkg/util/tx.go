package util

import (
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

// Message returns
func (tx *Tx) Message() *protocol.MsgTx {
	return tx.MsgTx
}
