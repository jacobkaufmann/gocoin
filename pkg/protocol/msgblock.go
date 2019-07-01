package protocol

import "io"

// MsgBlock transmits block data unsolicited or in response to a getdata message
// which requests block information from a block hash.
type MsgBlock struct {
	*BlockHeader
	Txns []*MsgTx
}

// NewMsgBlock returns a new block message.
func NewMsgBlock(hdr *BlockHeader, txns []*MsgTx) *MsgBlock {
	return &MsgBlock{
		BlockHeader: hdr,
		Txns:        txns,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgBlock) Serialize(w io.Writer, pver uint32) error {
	err := msg.BlockHeader.Serialize(w, pver)
	if err != nil {
		return err
	}

	for _, tx := range msg.Txns {
		err = tx.Serialize(w, pver)
		if err != nil {
			return err
		}
	}

	return nil
}

// Deserialize deserializes data from r into msg.
func (msg *MsgBlock) Deserialize(r io.Reader, pver uint32) error {
	err := msg.BlockHeader.Deserialize(r, pver)
	if err != nil {
		return err
	}

	for i := 0; i < int(msg.TxCount()); i++ {
		tx := &MsgTx{}
		tx.Deserialize(r, pver)
		if err != nil {
			return err
		}
		msg.Txns = append(msg.Txns, tx)
	}

	return nil
}

// TxCount returns the number of transactions in the block message.
func (msg *MsgBlock) TxCount() uint64 {
	return msg.TxCount()
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
