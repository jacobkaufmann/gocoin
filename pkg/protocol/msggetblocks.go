package protocol

import "io"

// MsgGetBlocks requests an inventory message that provides block header hashes
// starting from a particular point in the block chain.
type MsgGetBlocks struct {
	Version      uint32
	HeaderHashes []*[HashSize]byte
	StopHash     *[HashSize]byte
}

// NewMsgGetBlocks returns a new getblocks message for a set of header hashes
// and a stop hash.
func NewMsgGetBlocks(version uint32, hashes []*[HashSize]byte,
	stopHash *[HashSize]byte) *MsgGetBlocks {
	return &MsgGetBlocks{
		Version:      version,
		HeaderHashes: hashes,
		StopHash:     stopHash,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgGetBlocks) Serialize(w io.Writer, pver uint32) error {
	err := writeElement(w, msg.Version)
	if err != nil {
		return err
	}

	err = writeCompactSize(w, pver, msg.HashCount())
	if err != nil {
		return err
	}

	for _, hash := range msg.HeaderHashes {
		err = writeElement(w, hash)
		if err != nil {
			return err
		}
	}

	return writeElement(w, msg.StopHash)
}

// Deserialize deserializes data from r into msg.
func (msg *MsgGetBlocks) Deserialize(r io.Reader, pver uint32) error {
	err := readElement(r, &msg.Version)
	if err != nil {
		return err
	}

	var n uint64
	err = readCompactSize(r, pver, &n)
	if err != nil {
		return err
	}

	for i := 0; i < int(n); i++ {
		var hash [HashSize]byte
		err = readElement(r, &hash)
		if err != nil {
			return err
		}
		msg.HeaderHashes = append(msg.HeaderHashes, &hash)
	}

	return readElement(r, msg.StopHash)
}

// HashCount returns the number of block header hashes in the getblocks
// message.
func (msg *MsgGetBlocks) HashCount() uint64 {
	return uint64(len(msg.HeaderHashes))
}

// Command returns the message type of the getblocks message.
func (msg *MsgGetBlocks) Command() MsgType {
	return MsgTypeGetBlocks
}

// MaxPayloadSize returns the maximum size in bytes of the getblocks message.
func (msg *MsgGetBlocks) MaxPayloadSize(pver uint32) uint32 {
	return MaxMsgSize
}
