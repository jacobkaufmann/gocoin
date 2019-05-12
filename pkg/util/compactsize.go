package util

import (
	"encoding/binary"
)

// CompactSize represents a variable-length integer to indicate the number
// of bytes in a following piece of data.
type CompactSize uint64

// littleEndian is a convenience variable for binary.littleEndian.
var littleEndian = binary.LittleEndian

// Bytes returns a variable-length byte slice containing a prefix identifier
// and the integer encoded in little endian order.
func (c CompactSize) Bytes() []byte {
	switch v := uint64(c); {
	case v < 0xFD:
		b8 := make([]byte, 1, 1)
		b8 = append(b8, uint8(v))
		return b8
	case v <= 0xFFFF:
		b16 := make([]byte, 3, 3)
		b16[0] = 0xFD
		littleEndian.PutUint16(b16[1:], uint16(v))
		return b16
	case v <= 0xFFFFFFFF:
		b32 := make([]byte, 5, 5)
		b32[0] = 0xFE
		littleEndian.PutUint32(b32[1:], uint32(v))
		return b32
	default:
		b64 := make([]byte, 9, 9)
		b64[0] = 0xFF
		littleEndian.PutUint64(b64[1:], v)
		return b64
	}
}

// Value returns the uint64 value of the CompactSize.
func (c CompactSize) Value() uint64 {
	return uint64(c)
}
