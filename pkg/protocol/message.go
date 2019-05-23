package protocol

import (
	"crypto/sha256"
	"io"
)

const (
	// MagicSize is the size in bytes of the magic value in the message header.
	MagicSize = 4

	// CommandSize is the size in bytes of a command in the message header.
	// Commands are padded by nulls (0x00) to pad out the size.
	// Example: 'version\0\0\0\0\0'.
	CommandSize = 12

	// MessageSizeSize is the size in bytes of the size value in the message
	// header.
	MessageSizeSize = 4

	// ChecksumSize is the size in bytes of a message checksum.
	ChecksumSize = 4

	// MessageHeaderSize is the size in bytes of a message header.  It is the
	// sum of the Bitcoin network (magic) 4 bytes + the command 12 bytes +
	// the payload size 4 bytes + the checksum 4 bytes.
	MessageHeaderSize = MagicSize + CommandSize + MessageSizeSize + ChecksumSize
)

// MsgType defines a bitcoin protocol message type.
type MsgType string

// Bitcoin protocol message types.
const (
	MsgTypeVersion     MsgType = "version"
	MsgTypeVerAck      MsgType = "verack"
	MsgTypeAddr        MsgType = "addr"
	MsgTypeInv         MsgType = "inv"
	MsgTypeGetData     MsgType = "getdata"
	MsgTypeGetBlocks   MsgType = "getblocks"
	MsgTypeGetHeaders  MsgType = "getheaders"
	MsgTypeNotFound    MsgType = "notfound"
	MsgTypeTx          MsgType = "tx"
	MsgTypeBlock       MsgType = "block"
	MsgTypeHeaders     MsgType = "headers"
	MsgTypeGetAddr     MsgType = "getaddr"
	MsgTypeMempool     MsgType = "mempool"
	MsgTypePing        MsgType = "ping"
	MsgTypePong        MsgType = "pong"
	MsgTypeReject      MsgType = "reject"
	MsgTypeSendHeaders MsgType = "sendheaders"
)

func (msgType MsgType) String() string {
	return string(msgType)
}

// Bytes returns the message type encoded into a byte array.
func (msgType MsgType) Bytes() [CommandSize]byte {
	var b [CommandSize]byte
	s := msgType.String()
	copy(b[:], []byte(s))
	return b
}

// Serializable describes a Bitcoin serializable object.
type Serializable interface {
	// Serialize converts the object to bytes and writes it to a writer.
	Serialize(io.Writer, uint32) error

	// Deserialize reads bytes from a reader and converts those bytes into
	// an object.
	Deserialize(io.Reader, uint32) error
}

// Message represents a message sent over a TCP connection to a node in the
// bitcoin peer-to-peer network.
type Message interface {
	// A Message is a serializable Bitcoin object.
	Serializable

	// Command returns the command name for a message.  The command name
	// identifies the message type contained in the payload of a message.
	Command() MsgType

	// MaxPayloadLength returns the maximum length the message payload can be.
	MaxPayloadLength(uint32) uint32
}

// MessageHeader represents a message header in the Bitcoin network protocol.
type MessageHeader struct {
	magic    BitcoinNet
	command  MsgType
	size     uint32
	checksum [ChecksumSize]byte
}

// Magic returns the magic value of the message the message header corresponds
// to.
//
// The magic value indicates the messsage origin network.
func (hdr *MessageHeader) Magic() BitcoinNet {
	return hdr.magic
}

// Command returns the message type of the message the message header
// corresponds to.
func (hdr *MessageHeader) Command() MsgType {
	return hdr.command
}

// Size returns the size of the message the message header corresponds to.
func (hdr *MessageHeader) Size() uint32 {
	return hdr.Size()
}

// Checksum returns the checksum of the message the message header corresponds
// to.
func (hdr *MessageHeader) Checksum() [ChecksumSize]byte {
	return hdr.checksum
}

// checksum returns the prefix (of length ChecksumSize) of the double SHA256
// checksum of the data.
//
// The double SHA256 checksum is computed as SHA256(SHA256(data)).
func checksum(b []byte) (check [ChecksumSize]byte) {
	h := sha256.Sum256(b)
	dh := sha256.Sum256(h[:])
	copy(check[:], dh[:ChecksumSize])
	return
}

// WriteMessage writes a Message to a writer and returns the number of bytes
// written.
// func WriteMessage(w io.Writer, msg Message, pver uint32, net BitcoinNet) (int, error) {
//
// }

// ReadMessage reads and validates bytes from a reader and assembles a Message
// from those bytes.
// func ReadMessage(r io.Reader, pver uint32, net BitcoinNet) (Message, error) {
//
// }
