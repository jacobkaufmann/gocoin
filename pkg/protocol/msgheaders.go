package protocol

import (
	"io"
	"time"
)

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

// Serialize serializes msg and writes to w.
func (msg *MsgHeaders) Serialize(w io.Writer, pver uint32) error {
	count := CompactSize(uint64(len(msg.Headers)))
	err := count.Serialize(w, pver)
	if err != nil {
		return err
	}

	for _, v := range msg.Headers {
		err = v.Serialize(w, pver)
		if err != nil {
			return err
		}
	}

	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgHeaders) Deserialize(r io.Reader, pver uint32) error {
	var count CompactSize
	err := count.Deserialize(r, pver)
	if err != nil {
		return err
	}

	n := count.Uint64()
	for i := 0; i < n; i++ {
		hdr := &BlockHeader{}
		err = hdr.Deserialize(r, pver)
		if err != nil {
			return err
		}

		msg.Headers = append(msg.Headers, hdr)
	}

	return nil
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
	Time           time.Time
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
		Time:           time.Unix(time.Now().Unix(), 0),
		NumBits:        numBits,
		Nonce:          nonce,
		TxCount:        txCount,
	}
}

// Serialize serializes hdr and writes to w.
func (hdr *BlockHeader) Serialize(w io.Writer, pver uint32) error {
	err := writeElements(w, hdr.Version, hdr.PrevBlockHash, hdr.MerkleRootHash,
		hdr.time.Unix(), hdr.NumBits, hdr.Nonce)
	if err != nil {
		return err
	}
	return hdr.TxCount.Serialize(w, pver)
}

// Deserialize deserializes data from r into hdr.
func (hdr *BlockHeader) Deserialize(r io.Reader, pver uint32) error {
	var time int64timestamp
	err := readElements(r, &hdr.Version, &hdr.PrevBlockHash,
		&hdr.MerkleRootHash, &time, &hdr.NumBits, &hdr.Nonce)
	if err != nil {
		return err
	}

	hdr.Time = time.Unix(time, 0)
	return hdr.TxCount.Deserialize(r, pver)
}
