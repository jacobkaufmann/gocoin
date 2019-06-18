package protocol

import "io"

// CompactSize represents a variable-length integer to indicate the number
// of bytes in a following piece of data.
type CompactSize uint64

var prefixes map[byte]uint32

func init() {
	prefixes[0xFD] = 3
	prefixes[0xFE] = 5
	prefixes[0xFF] = 9
}

// Serialize serializes c and writes to w.
func (c CompactSize) Serialize(w io.Writer, pver uint32) error {
	b := c.bytes(pver)
	n, err := w.Write(b)
	if err != nil {
		return err
	}
	if uint32(n) != c.Size() {
		return ErrInsufficientBytesWritten
	}

	return nil
}

// Deserialize deserializes data from r into c.
func (c CompactSize) Deserialize(r io.Reader, pver uint32) error {
	var prefix [1]byte
	_, err := r.Read(prefix[:])
	if err != nil {
		return err
	}

	n, ok := prefixes[prefix[0]]
	if !ok {
		n = 1
	}

	b := make([]byte, n)
	if n == 1 {
		b[0] = prefix[0]
	}
	c = compactSizeFromBytes(prefix[0], b)
	return nil
}

// compactSizeFromBytes parses a little endian ordered byte slice and returns
// a CompactSize object.
func compactSizeFromBytes(prefix byte, b []byte) CompactSize {
	switch prefix {
	case 0xFD:
		return CompactSize(uint64(littleEndian.Uint16(b)))
	case 0xFE:
		return CompactSize(uint64(littleEndian.Uint32(b)))
	case 0xFF:
		return CompactSize(uint64(littleEndian.Uint64(b)))
	default:
		return CompactSize(uint64(uint8(b[0])))
	}
}

// bytes returns a variable-length byte slice containing a prefix identifier
// and the integer encoded in little endian order.
func (c CompactSize) bytes(pver uint32) []byte {
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

// Uint64 returns the uint64 value of the CompactSize.
func (c CompactSize) Uint64() uint64 {
	return uint64(c)
}

// Size returns the length of the byte representation of the CompactSize.
func (c CompactSize) Size() uint32 {
	switch v := uint64(c); {
	case v < 0xFD:
		return 1
	case v <= 0xFFFF:
		return 3
	case v <= 0xFFFFFFFF:
		return 5
	default:
		return 9
	}
}
