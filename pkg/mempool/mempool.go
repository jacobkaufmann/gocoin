package mempool

import (
	"sync"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
	"github.com/jacobkaufmann/gocoin/pkg/util"
)

// MemPool represents the transaction mempool, which holds valid but
// unconfirmed transactions. MemPool implements the TxPool interface.
type MemPool struct {
	txns map[hashing.Hash]*Entry
	*sync.RWMutex
}

// Entry stores data about the corresponding transaction as well as
// data about all in-mempool transactions that depend on the transactions,
// or its "descendants".
type Entry struct {
	Tx                   *util.Tx
	Fee                  util.Amount
	CountWithDescendants uint64
	CountWithAncestors   uint64
}

// newEntry returns a new mempool entry for tx.
func newEntry(tx *util.Tx) *Entry {
	return &Entry{
		Tx: tx,
	}
}

// New returns a new mempool.
func New() *MemPool {
	txns := make(map[hashing.Hash]*Entry)
	return &MemPool{txns}
}

// Get returns a transaction entry specified by id if it exists in the mempool.
func (mp *MemPool) Get(id *hashing.Hash) *Entry {
	mp.RLock()
	defer mp.RUnlock()

	tx, ok := mp.txns[id]
	if ok == false {
		return nil
	}
	return tx
}

// Insert adds a transaction entry to the mempool if it does not exist.
func (mp *MemPool) Insert(tx *util.Tx, pver uint32) {
	mp.Lock()
	defer mp.Unlock()

	entry := newEntry(tx)

	// TODO: update fields of entry for related transactions in mempool.

	mp.txns[tx.TxID(pver)] = entry
}

// Remove removes a transaction entry specified by id if it exists in the
// mempool.
func (mp *MemPool) Remove(id *hashing.Hash) bool {
	mp.Lock()
	defer mp.Unlock()

	_, ok := mp.txns[id]
	if ok == false {
		return false
	}

	delete(mp.txns, id)
	return true
}

// Clear removes all transactions from the mempool.
func (mp *MemPool) Clear() {
	mp.Lock()
	defer mp.Unlock()

	mp.txns = make(map[hashing.Hash]*util.Tx)
}
