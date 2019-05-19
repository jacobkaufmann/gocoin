package protocol

// HashSize represents the size (in bytes) of a SHA256 checksum.
const HashSize = 32

// BitcoinNet is the magic value associated with a Bitcoin network.
type BitcoinNet uint32

const (
	// MainNet is the main Bitcoin network.
	MainNet BitcoinNet = 0xD9B4BEF9

	// TestNet is the test Bitcoin network.
	TestNet BitcoinNet = 0xDAB5BFFA
)

const (
	// MaxInvSize is the maximum number of entries in an inventory protocol
	// message.
	MaxInvSize = 50000

	// MaxMsgLength is the maximum length of incoming protocol messages.
	MaxMsgLength = 4 * 1000 * 1000

	// MaxAddrToSend is the maximum number of new addresses to accumulate
	// before announcing.
	MaxAddrToSend = 1000
)
