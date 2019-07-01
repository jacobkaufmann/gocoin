package protocol

import "io"

// A MsgTx transmits a single bitcoin transaction.
type MsgTx struct {
	Version  int32
	Inputs   []*TxIn
	Outputs  []*TxOut
	LockTime uint32
}

// NewMsgTx returns a new transaction message for a set of inputs and a set of
// outputs and a specified lock time.
func NewMsgTx(version int32, inputs []*TxIn, outputs []*TxOut, lockTime uint32) *MsgTx {
	return &MsgTx{
		Version:  version,
		Inputs:   inputs,
		Outputs:  outputs,
		LockTime: lockTime,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgTx) Serialize(w io.Writer, pver uint32) error {
	err := writeElement(w, msg.Version)
	if err != nil {
		return err
	}

	// Transaction inputs.
	err = writeCompactSize(w, pver, msg.TxInCount())
	if err != nil {
		return err
	}
	for _, input := range msg.Inputs {
		err = input.Serialize(w, pver)
		if err != nil {
			return err
		}
	}

	// Transaction outputs.
	err = writeCompactSize(w, pver, msg.TxOutCount())
	if err != nil {
		return err
	}
	for _, output := range msg.Outputs {
		err = output.Serialize(w, pver)
		if err != nil {
			return err
		}
	}

	return writeElement(w, msg.LockTime)
}

// Deserialize deserializes data from r into msg.
func (msg *MsgTx) Deserialize(r io.Reader, pver uint32) error {
	err := readElement(r, &msg.Version)
	if err != nil {
		return err
	}

	var n uint64

	// Transaction inputs.
	err = readCompactSize(r, pver, &n)
	if err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		input := &TxIn{}
		err = input.Deserialize(r, pver)
		if err != nil {
			return err
		}
		msg.Inputs = append(msg.Inputs, input)
	}

	// Transaction outputs.
	err = readCompactSize(r, pver, &n)
	if err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		output := &TxOut{}
		err = output.Deserialize(r, pver)
		if err != nil {
			return err
		}
		msg.Outputs = append(msg.Outputs, output)
	}

	return readElement(r, &msg.LockTime)
}

// TxInCount returns the number of transaction inputs.
func (msg *MsgTx) TxInCount() uint64 {
	return uint64(len(msg.Inputs))
}

// TxOutCount returns the number of transaction outputs
func (msg *MsgTx) TxOutCount() uint64 {
	return uint64(len(msg.Outputs))
}

// Command returns the message type of the transaction message.
func (msg *MsgTx) Command() MsgType {
	return MsgTypeTx
}

// MaxPayloadSize returns the maximum size in bytes of the transaction
// message.
func (msg *MsgTx) MaxPayloadSize(pver uint32) uint32 {
	return MaxMsgSize
}

// A TxIn is an input to a transaction.
type TxIn struct {
	PrevOutput       TxOutPoint
	ScriptUnlockSize uint64
	ScriptUnlock     []byte
	Sequence         uint32
}

// Serialize serializes in and writes to w.
func (in *TxIn) Serialize(w io.Writer, pver uint32) error {
	err := in.PrevOutput.Serialize(w, pver)
	if err != nil {
		return err
	}

	err = writeCompactSize(w, pver, in.ScriptUnlockSize)
	if err != nil {
		return err
	}

	return writeElements(w, in.ScriptUnlock, in.Sequence)
}

// Deserialize deserializes data from r into in.
func (in *TxIn) Deserialize(r io.Reader, pver uint32) error {
	err := in.PrevOutput.Deserialize(r, pver)
	if err != nil {
		return err
	}

	err = readCompactSize(r, pver, &in.ScriptUnlockSize)
	if err != nil {
		return err
	}

	return readElements(r, in.ScriptUnlock, &in.Sequence)
}

// A TxOut is an output of a transaction.
type TxOut struct {
	Value          int64
	ScriptLockSize uint64
	ScriptLock     []byte
}

// Serialize serializes out and writes to w.
func (out *TxOut) Serialize(w io.Writer, pver uint32) error {
	err := writeElement(w, out.Value)
	if err != nil {
		return err
	}

	err = writeCompactSize(w, pver, out.ScriptLockSize)
	if err != nil {
		return err
	}

	return writeElement(w, out.ScriptLock)
}

// Deserialize deserializes data from r into out.
func (out *TxOut) Deserialize(r io.Reader, pver uint32) error {
	err := readElement(r, &out.Value)
	if err != nil {
		return err
	}

	err = readCompactSize(r, pver, &out.ScriptLockSize)
	if err != nil {
		return err
	}

	return readElement(r, &out.ScriptLock)
}

// A TxOutPoint contains information to refer to a specific transaction output.
type TxOutPoint struct {
	Hash  *[HashSize]byte
	Index uint32
}

// txOutPointSize is the size of the byte representation of a TxOutPoint.
const txOutPointSize = HashSize + 4

// Serialize serializes outPoint and writes to w.
func (outPoint *TxOutPoint) Serialize(w io.Writer, pver uint32) error {
	return writeElements(w, outPoint.Hash, outPoint.Index)
}

// Deserialize deserializes data from r into outPoint.
func (outPoint *TxOutPoint) Deserialize(r io.Reader, pver uint32) error {
	return readElements(r, &outPoint.Hash, &outPoint.Index)
}
