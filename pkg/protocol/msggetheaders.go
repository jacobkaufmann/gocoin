package protocol

// MsgGetHeaders requests a headers message containing block headers starting
// from a particular point in the block chain.
type MsgGetHeaders struct {
	Version      uint32
	HeaderHashes []*[HashSize]byte
	StopHash     *[HashSize]byte
}

// HashCount returns the number of block header hashes in the getheaders
// message.
func (msg *MsgGetHeaders) HashCount() CompactSize {
	return CompactSize(len(msg.HeaderHashes))
}

// Command returns the message type of the getheaders message.
func (msg *MsgGetHeaders) Command() MsgType {
	return MsgTypeGetHeaders
}

// MaxPayloadSize returns the maximum size in bytes of the getheaders message.
func (msg *MsgGetHeaders) MaxPayloadSize(pver uint32) uint32 {
	return MaxMsgSize
}
