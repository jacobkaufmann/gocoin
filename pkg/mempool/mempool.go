package mempool

import (
	"sync"

	"github.com/jacobkaufmann/gocoin/pkg/crypto/hashing"
	"github.com/jacobkaufmann/gocoin/pkg/util"
)

// MemPool represents the transaction mempool, which holds valid but
// unconfirmed transactions. MemPool implements the TxPool interface.
type MemPool struct {
	txns map[hashing.Hash]*util.Tx
	*sync.RWMutex
}

// New returns a new mempool.
func New() *MemPool {
	txns := make(map[hashing.Hash]*util.Tx)
	return &MemPool{txns}
}

// Get returns a transaction specified by id if it exists in the mempool.
func (mp *MemPool) Get(id *hashing.Hash) *util.Tx {
	mp.RLock()
	defer mp.RUnlock()

	tx, ok := mp.txns[id]
	if ok == false {
		return nil
	}
	return tx
}

// Insert adds a transaction to the mempool if it does not exist.
func (mp *MemPool) Insert(tx *util.Tx, pver uint32) {
	mp.Lock()
	defer mp.Unlock()

	mp.txns[tx.TxID(pver)] = tx
}

// Remove removes a transaction specified by id if it exists in the mempool.
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
