package mempool

import (
	"github.com/jacobkaufmann/gocoin/pkg/protocol"
)

// A UTXO represents an unspent transaction output.  Unspent transaction
// outputs are the outputs of previous transactions which are eligible to
// be consumed by future transactions.
type UTXO struct {
	*protocol.TxOut
}
