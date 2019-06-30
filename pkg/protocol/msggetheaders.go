package protocol

import "io"

// MsgGetHeaders requests a headers message containing block headers starting
// from a particular point in the block chain.
type MsgGetHeaders struct {
	Version      uint32
	HeaderHashes []*[HashSize]byte
	StopHash     *[HashSize]byte
}

// NewMsgGetHeaders returns a new getheaders message for a set of header hashes
// and a stop hash.
func NewMsgGetHeaders(version uint32, hashes []*[HashSize]byte,
	stopHash *[HashSize]byte) *MsgGetHeaders {

	return &MsgGetHeaders{
		Version:      version,
		HeaderHashes: hashes,
		StopHash:     stopHash,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgGetHeaders) Serialize(w io.Writer, pver uint32) error {
	err := writeElement(w, msg.Version)
	if err != nil {
		return err
	}

	err = writeCompactSize(w, pver, msg.HashCount())
	if err != nil {
		return err
	}

	for _, v := range msg.HeaderHashes {
		err = writeElement(w, v)
		if err != nil {
			return err
		}
	}

	return writeElement(w, msg.StopHash)
}

// Deserialize deserializes data from r into msg.
func (msg *MsgGetHeaders) Deserialize(r io.Reader, pver uint32) error {
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
		msg.HeaderHashes = append(msg.HeaderHashes, hash)
	}

	return readElement(r, msg.StopHash)
}

// HashCount returns the number of block header hashes in the getheaders
// message.
func (msg *MsgGetHeaders) HashCount() uint64 {
	return len(msg.HeaderHashes)
}

// Command returns the message type of the getheaders message.
func (msg *MsgGetHeaders) Command() MsgType {
	return MsgTypeGetHeaders
}

// MaxPayloadSize returns the maximum size in bytes of the getheaders message.
func (msg *MsgGetHeaders) MaxPayloadSize(pver uint32) uint32 {
	return MaxMsgSize
}
