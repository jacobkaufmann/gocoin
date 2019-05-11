package protocol

// BitcoinNet is the magic value associated with a Bitcoin network.
type BitcoinNet uint32

const (
	// MainNet is the main Bitcoin network.
	MainNet BitcoinNet = 0xD9B4BEF9

	// TestNet is the test Bitcoin network.
	TestNet BitcoinNet = 0xDAB5BFFA
)
