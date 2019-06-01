package protocol

import (
	"encoding/binary"
)

var (
	// littleEndian is a convenience variable for binary.littleEndian.
	littleEndian = binary.LittleEndian

	// bigEndian is a convenience variable for binary.BigEndian.
	bigEndian = binary.BigEndian
)

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
