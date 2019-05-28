package protocol

// A ServiceFlag is a bit flag used to identify whether a particular feature
// is to be supported in a connection between nodes.
type ServiceFlag uint64

// Bitcoin service flags.
const (
	SFNetwork        ServiceFlag = 1
	SFGetUTxO        ServiceFlag = 2
	SFBloom          ServiceFlag = 4
	SFWitness        ServiceFlag = 8
	SFNetworkLimited ServiceFlag = 1024
)
