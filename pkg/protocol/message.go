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

	// ChecksumSize is the size in bytes of the message checksum in the message
	// header.
	ChecksumSize = 4

	// MessageSizeOffset is the starting position of the messsage size in the
	// message header.
	MessageSizeOffset = MagicSize + CommandSize

	// ChecksumOffset is the starting position of the message checksum in the
	// message header.
	ChecksumOffset = MessageSizeOffset + ChecksumSize

	// MessageHeaderSize is the size in bytes of a message header.  It is the
	// sum of the Bitcoin network (magic) 4 bytes + the command 12 bytes +
	// the payload size 4 bytes + the checksum 4 bytes.
	MessageHeaderSize = MagicSize + CommandSize + MessageSizeSize + ChecksumSize

	// EmptyPayloadSize is a convenience variable for messages which have no
	// payload.
	EmptyPayloadSize = 0
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
	// Serialize converts an object into a byte format specified by the
	// protocol version and writes those bytes to a writer.
	Serialize(io.Writer, uint32) error

	// Deserialize reads bytes from a reader and converts those bytes into
	// an object according to the specified protocol version.
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

	// MaxPayloadSize returns the maximum size in bytes the message payload can
	// be.
	MaxPayloadSize(uint32) uint32
}

// messageHeader represents a message header in the bitcoin network protocol.
type messageHeader struct {
	magic    BitcoinNet
	command  MsgType
	size     uint32
	checksum [ChecksumSize]byte
}

// readMessageHeader reads from r and returns a message header.
func readMessageHeader(r io.Reader) (*messageHeader, error) {
	buf := make([]byte, MessageHeaderSize)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}

	magic := buf[:MagicSize]
	cmd := buf[MagicSize:MessageSizeOffset]
	size := buf[MessageSizeOffset:ChecksumOffset]

	var check [ChecksumSize]byte
	copy(check[:], buf[ChecksumOffset:])
	reverseBytes(check[:])

	hdr := &messageHeader{
		magic:    BitcoinNet(littleEndian.Uint32(magic)),
		command:  MsgType(cmd),
		size:     littleEndian.Uint32(size),
		checksum: check,
	}
	return hdr, nil
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
