package util

// Amount represents a bitcoin amount in satoshis (can be negative).
type Amount int64

const (
	// Coin is the number of satoshis in a Bitcoin.
	Coin Amount = 100000000

	// MaxCoins is the maximum number of satoshis that will ever be allocated.
	MaxCoins Amount = 21000000 * Coin
)
