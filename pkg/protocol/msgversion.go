package protocol

import "time"

// MsgVersion provides information about the transmitting node to the receiving
// node upon creating an outgoing connection. Until both peers have exchanged
// version messages, no other messages will be accepted.
//
// If the version message is accepted, the receiving node should send a verack
// message.
type MsgVersion struct {
	Version     uint32
	Services    ServiceFlag
	Timestamp   time.Time
	AddrRecv    *NetAddress
	AddrFrom    *NetAddress
	Nonce       uint64
	UserAgent   VarStr
	StartHeight uint32
}

// Command returns the message type of the version message.
func (msg *MsgVersion) Command() MsgType {
	return MsgTypeVersion
}

// MaxPayloadSize returns the maximum size in bytes of the version message.
func (msg *MsgVersion) MaxPayloadSize(pver uint32) uint32 {
	return 48
}
