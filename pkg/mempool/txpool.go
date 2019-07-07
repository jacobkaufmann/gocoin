package mempool

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
	"github.com/jacobkaufmann/gocoin/pkg/util"
)

// TxPool represents a pool of transactions. It provides methods for retrieving
// a transaction as well as for inserting and removing transactions.
type TxPool interface {
	Get(*hashing.Hash) *util.Tx
	Insert(*util.Tx)
	Remove(*hashing.Hash) bool
	Clear()
}
