package protocol

// MsgBlock transmits block data unsolicited or in response to a getdata message
// which requests block information from a block hash.
type MsgBlock struct {
	*BlockHeader
	Txns []*MsgTx
}

// Command returns the message type of the block message.
func (msg *MsgBlock) Command() MsgType {
	return MsgTypeBlock
}

// MaxPayloadSize returns the maximum size in bytes of the block message.
func (msg *MsgBlock) MaxPayloadSize(pver uint32) uint32 {
	// 1 MB
	return 1000 * 1000
}
