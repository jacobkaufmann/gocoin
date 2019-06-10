package protocol

// MsgReject informs the receiving node that one of its previous messages have
// been rejected.
type MsgReject struct {
	Msg    VarStr
	Code   byte
	Reason VarStr
	Data   *[HashSize]byte
}

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

// Command returns the message type of the reject message.
func (msg *MsgReject) Command() MsgType {
	return MsgTypeReject
}

// MaxPayloadSize returns the maximum size in bytes of the reject message.
// func (msg *MsgReject) MaxPayloadSize(pver uint32) uint32 {
//
// }
