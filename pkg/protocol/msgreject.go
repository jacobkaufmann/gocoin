package protocol

import "io"

// A RejectCode indicates why a message was rejected.
type RejectCode byte

// Bitcoin reject codes.
const (
	RejectCodeMalformed       RejectCode = 0x01
	RejectCodeInvalid         RejectCode = 0x10
	RejectCodeObsolete        RejectCode = 0x11
	RejectCodeDuplicate       RejectCode = 0x12
	RejectCodeNonstandard     RejectCode = 0x40
	RejectCodeDust            RejectCode = 0x41
	RejectCodeInsufficientFee RejectCode = 0x42
	RejectCodeCheckpoint      RejectCode = 0x43
)

// MsgReject informs the receiving node that one of its previous messages have
// been rejected.
type MsgReject struct {
	Msg    string
	Code   RejectCode
	Reason string
	Data   *[HashSize]byte
}

// NewMsgReject returns a new reject message.
func NewMsgReject(msg string, code RejectCode, reason string,
	data *[HashSize]byte) *MsgReject {
	return &MsgReject{
		Msg:    msg,
		Code:   code,
		Reason: reason,
		Data:   data,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgReject) Serialize(w io.Writer, pver uint32) error {
	err := writeVarStr(w, pver, msg.Msg)
	if err != nil {
		return err
	}

	err = writeElement(w, msg.Code)
	if err != nil {
		return err
	}

	err = writeVarStr(w, pver, msg.Reason)
	if err != nil {
		return err
	}

	return writeElement(w, msg.Data)
}

// Deserialize deserializes data from r into msg.
func (msg *MsgReject) Deserialize(r io.Reader, pver uint32) error {
	err := readVarStr(r, pver, &msg.Msg)
	if err != nil {
		return err
	}

	err = readElement(r, &msg.Code)
	if err != nil {
		return err
	}

	err = readVarStr(r, pver, &msg.Reason)
	if err != nil {
		return err
	}

	return readElement(r, &msg.Data)
}

// Command returns the message type of the reject message.
func (msg *MsgReject) Command() MsgType {
	return MsgTypeReject
}

// MaxPayloadSize returns the maximum size in bytes of the reject message.
func (msg *MsgReject) MaxPayloadSize(pver uint32) uint32 {
	return MaxMsgSize
}
