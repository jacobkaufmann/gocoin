package protocol

// MsgGetBlocks requests an inventory message that provides block header hashes
// starting from a particular point in the block chain.
type MsgGetBlocks struct {
	Version      uint32
	HeaderHashes []*[HashSize]byte
	StopHash     *[HashSize]byte
}

// HashCount returns the number of block header hashes in the getblocks
// message.
func (msg *MsgGetBlocks) HashCount() CompactSize {
	return CompactSize(len(msg.HeaderHashes))
}

// Command returns the message type of the getblocks message.
func (msg *MsgGetBlocks) Command() MsgType {
	return MsgTypeGetAddr
}

// MaxPayloadSize returns the maximum size in bytes of the getblocks message.
func (msg *MsgGetBlocks) MaxPayloadSize(pver uint32) uint32 {
	return MaxMsgSize
}
