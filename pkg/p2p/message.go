package p2p

// messageHeader represents a message header in the Bitcoin network protocol.
type messageHeader struct {
	magic    uint32
	cmd      command
	size     uint32
	checksum uint32
}

type command [12]byte

type messageType string

const (
	version     messageType = "version"
	verack      messageType = "verack"
	addr        messageType = "addr"
	inv         messageType = "inv"
	getData     messageType = "getdata"
	getBlocks   messageType = "getblocks"
	getHeaders  messageType = "getheaders"
	notFound    messageType = "notfound"
	tx          messageType = "tx"
	block       messageType = "block"
	headers     messageType = "headers"
	getAddr     messageType = "getaddr"
	mempool     messageType = "mempool"
	ping        messageType = "ping"
	pong        messageType = "pong"
	reject      messageType = "reject"
	sendHeaders messageType = "sendheaders"
)

func (t *messageType) cmd() command {
	var cmd command
	copy(cmd[:], string(*t))
	return cmd
}
