package protocol

// messageHeader represents a message header in the Bitcoin network protocol.
type messageHeader struct {
	magic    uint32
	command  string
	size     uint32
	checksum uint32
}

// MessageHeaderSize is the size in bytes of a message header.  It is the
// sum of the Bitcoin network (magic) 4 bytes + the command 12 bytes +
// the payload size 4 bytes + the checksum 4 bytes.
const MessageHeaderSize = 24

// CommandSize is the size in bytes of a command in the message header.
// Commands are padded by nulls (0x00) to pad out the size.
// Example: 'version\0\0\0\0\0'.
const CommandSize = 12

const (
	CmdVersion     = "version"
	CmdVerAck      = "verack"
	CmdAddr        = "addr"
	CmdInv         = "inv"
	CmdGetData     = "getdata"
	CmdGetBlocks   = "getblocks"
	CmdGetHeaders  = "getheaders"
	CmdNotFound    = "notfound"
	CmdTx          = "tx"
	CmdBlock       = "block"
	CmdHeaders     = "headers"
	CmdGetAddr     = "getaddr"
	CmdMempool     = "mempool"
	CmdPing        = "ping"
	CmdPong        = "pong"
	CmdReject      = "reject"
	CmdSendHeaders = "sendheaders"
)

// Message represents a message sent over a TCP connection to a node in the
// Bitcoin peer-to-peer network.
type Message interface {
	// Encode encodes a message into a destination byte slice.
	Encode(dst []byte, version uint32) error

	// Decode decodes a source byte slice into a Message.
	// TODO: evaluate semantics
	Decode(src []byte, version uint32) error

	// Command returns the command name for a message.  The command name
	// identifies the message type contained in the payload of a message.
	Command() string
}
