package wallet

import (
	"github.com/jacobkaufmann/gocoin/pkg/crypto/btcec"
)

// Wallet represents a full-service bitcoin wallet capable of distributing
// keys and addresses and creating, signing, and broadcasting transactions.
type Wallet interface {
	Balance() uint64
	AvailableBalance() uint64
	PubKey() *btcec.PublicKey
}

// A BtcWallet represents a bitcoin wallet.
type BtcWallet struct {
	privKey *btcec.PrivateKey
}
