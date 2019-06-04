package protocol

// MsgHeaders transmits block headers in response to a getheaders message.
type MsgHeaders struct {
	Headers []*BlockHeader
}

const (
	// BlockHeaderSize is the size in bytes of a serialized block header.
	BlockHeaderSize = 80

	// MaxBlockHeaders is the maximum number of block headers capable of being
	// sent in a single message.
	MaxBlockHeaders = 2000
)

// NewMsgHeaders returns a new headers message containing headers.
func NewMsgHeaders(headers []*BlockHeader) *MsgHeaders {
	return &MsgHeaders{
		Headers: headers,
	}
}

// Count returns the number of headers in the headers message.
func (msg *MsgHeaders) Count() CompactSize {
	return CompactSize(len(msg.Headers))
}

// Command returns the message type of the headers message.
func (msg *MsgHeaders) Command() MsgType {
	return MsgTypeHeaders
}

// MaxPayloadSize returns the maximum size in bytes of the headers message.
func (msg *MsgHeaders) MaxPayloadSize(pver uint32) uint32 {
	// A suffix of 0x00 is included with each block header for the transaction
	// count. The transaction count is always zero because the headers message
	// doesn't include any transactions.
	return uint32(MaxBlockHeaders*(BlockHeaderSize+1) + 9)
}

// A BlockHeader contains block metadata.
type BlockHeader struct {
	Version        uint32
	PrevBlockHash  *[HashSize]byte
	MerkleRootHash *[HashSize]byte
	Time           uint32
	NumBits        uint32
	Nonce          uint32
	TxCount        CompactSize
}

// NewBlockHeader returns a new block header with the specified metadata.
func NewBlockHeader(version uint32, prevBlock, merkleRoot *[HashSize]byte,
	time, numBits, nonce uint32, txCount CompactSize) *BlockHeader {
	return &BlockHeader{
		Version:        version,
		PrevBlockHash:  prevBlock,
		MerkleRootHash: merkleRoot,
		Time:           time,
		NumBits:        numBits,
		Nonce:          nonce,
		TxCount:        txCount,
	}
}
