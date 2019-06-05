package protocol

// A MsgTx transmits a single bitcoin transaction.
type MsgTx struct {
	Version  int32
	Inputs   []*TxIn
	Outputs  []*TxOut
	LockTime uint32
}

// NewMsgTx returns a transaction message containing a set of inputs and a set
// of outputs and a specified lock time.
func NewMsgTx(version int32, inputs []*TxIn, outputs []*TxOut, lockTime uint32) *MsgTx {
	return &MsgTx{
		Version:  version,
		Inputs:   inputs,
		Outputs:  outputs,
		LockTime: lockTime,
	}
}

// TxInCount returns the number of transaction inputs.
func (msg *MsgTx) TxInCount() CompactSize {
	return CompactSize(len(msg.Inputs))
}

// TxOutCount returns the number of transaction outputs
func (msg *MsgTx) TxOutCount() CompactSize {
	return CompactSize(len(msg.Outputs))
}

// Size returns the size in bytes of the transaction message.
func (msg *MsgTx) Size() uint32 {
	var size = 8 + msg.TxInCount().Size() + msg.TxOutCount().Size()
	for i := 0; i < len(msg.Inputs); i++ {
		size += msg.Inputs[i].Size()
	}
	for i := 0; i < len(msg.Outputs); i++ {
		size += msg.Outputs[i].Size()
	}
	return size
}

// Command returns the message type of the transaction message.
func (msg *MsgTx) Command() MsgType {
	return MsgTypeTx
}

// MaxPayloadSize returns the maximum size in bytes of the transaction
// message.
// func (msg *MsgTx) MaxPayloadSize(pver uint32) uint32 {
//
// }

// A TxIn is an input to a transaction.
type TxIn struct {
	PrevOutput     TxOutPoint
	ScriptLockSize CompactSize
	ScriptLock     []byte
	Sequence       uint32
}

// Size returns the size of the byte representation of a TxIn.
func (in *TxIn) Size() uint32 {
	// TODO: find a more elegant/safe way to compute this without so many type
	// conversions.
	size := uint64(txOutPointSize)
	size += uint64(in.ScriptLockSize.Size()) + in.ScriptLockSize.Uint64()
	size += 4
	return uint32(size)
}

// A TxOut is an output of a transaction.
type TxOut struct {
	Value            uint64
	ScriptUnlockSize CompactSize
	ScriptUnlock     []byte
}

// Size returns the size of the byte representation of a TxOut.
func (out *TxOut) Size() uint32 {
	// TODO: find a more elegant/safe way to compute this without so many type
	// conversions.
	size := 8 + uint64(out.ScriptUnlockSize.Size()) + out.ScriptUnlockSize.Uint64()
	return uint32(size)
}

// A TxOutPoint contains information to refer to a specific transaction output.
type TxOutPoint struct {
	Hash  *[HashSize]byte
	Index uint32
}

const txOutPointSize uint32 = HashSize + 4

// Size returns the size of the byte representation for a TxOutPoint.
func (outPoint *TxOutPoint) Size() uint32 {
	return txOutPointSize
}
