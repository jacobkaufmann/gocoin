package protocol

import (
	"io"
	"time"
)

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
	UserAgent   string
	StartHeight uint32
}

// NewMsgVersion returns a new version message.
func NewMsgVersion(version uint32, addrRecv *NetAddress, addrFrom *NetAddress,
	nonce uint64, userAgent string, startHeight uint32) *MsgVersion {
	return &MsgVersion{
		Version:     version,
		Services:    0,
		Timestamp:   time.Unix(time.Now().Unix(), 0),
		AddrRecv:    addrRecv,
		AddrFrom:    addrFrom,
		Nonce:       nonce,
		UserAgent:   userAgent,
		StartHeight: startHeight,
	}
}

// Serialize serializes msg and writes to w.
func (msg *MsgVersion) Serialize(w io.Writer, pver uint32) error {
	err := writeElements(w, msg.Version, msg.Services, msg.Timestamp.Unix())
	if err != nil {
		return err
	}

	// Network addresses do not include timestamp in version messages.
	err = writeNetAddress(w, pver, *msg.AddrRecv, false)
	if err != nil {
		return err
	}
	err = writeNetAddress(w, pver, *msg.AddrFrom, false)
	if err != nil {
		return err
	}

	err = writeElement(w, msg.Nonce)
	if err != nil {
		return err
	}

	err = writeVarStr(w, pver, msg.UserAgent)
	if err != nil {
		return err
	}

	return writeElement(w, msg.StartHeight)
}

// Deserialize deserializes data from r into msg.
func (msg *MsgVersion) Deserialize(r io.Reader, pver uint32) error {
	var timestamp int64Time

	err := readElements(r, &msg.Version, &msg.Services, &timestamp)
	if err != nil {
		return err
	}
	msg.Timestamp = time.Time(timestamp)

	// Network addresses do not include timestamp in version messages.
	err = readNetAddress(r, pver, msg.AddrRecv, false)
	if err != nil {
		return err
	}
	err = readNetAddress(r, pver, msg.AddrFrom, false)
	if err != nil {
		return err
	}

	err = readElement(r, &msg.Nonce)
	if err != nil {
		return err
	}

	err = readVarStr(r, pver, &msg.UserAgent)
	if err != nil {
		return err
	}

	return readElement(r, &msg.StartHeight)
}

// Command returns the message type of the version message.
func (msg *MsgVersion) Command() MsgType {
	return MsgTypeVersion
}

// MaxPayloadSize returns the maximum size in bytes of the version message.
func (msg *MsgVersion) MaxPayloadSize(pver uint32) uint32 {
	return 48
}
