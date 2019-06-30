// Variable-length integer utilites for serialization and deserialization. A
// CompactSize is the variable-length integer encoding used by the Bitcoin
// protocol. A separate type is unnecessary because we only need the
// variable-length representation at serialization time.

package protocol

import "io"

// emptyPrefix indicates the non-existance of a prefix when a CompactSize
// may be represented by a single byte.
const emptyPrefix = 0x00

// writeCompactSize encodes a uint64 into a CompactSize and writes the
// representation to w.
func writeCompactSize(w io.Writer, pver uint32, val uint64) error {
	prefix, v := makeCompactSize(val, pver)
	b := make([]byte, compactSizeSize(prefix, pver))

	if prefix == emptyPrefix {
		b[0] = v.(uint8)
	} else {
		b[0] = prefix
		switch cmpct := v.(type) {
		case uint16:
			littleEndian.PutUint16(b[1:], cmpct)
		case uint32:
			littleEndian.PutUint32(b[1:], cmpct)
		case uint64:
			littleEndian.PutUint64(b[1:], cmpct)
		}
	}

	_, err := w.Write(b)
	if err != nil {
		return err
	}
	return nil
}

// readCompactSize reads from r and decodes the CompactSize representation
// into a uint64.
func readCompactSize(r io.Reader, pver uint32, val *uint64) error {
	prefix := make([]byte, 1)
	_, err := r.Read(prefix)
	if err != nil {
		return err
	}

	size := compactSizeSize(prefix[0], pver)
	b := make([]byte, size)
	_, err = r.Read(b)
	if err != nil {
		return err
	}

	*val = littleEndian.Uint64(b)
	return nil
}

// makeCompactSize returns the CompactSize representation (prefix and value)
// for a uint64. If the unsigned integer can be represented by a single byte,
// the prefix is returned as a zero value.
func makeCompactSize(val uint64, pver uint32) (prefix byte, v interface{}) {
	switch {
	case val < 0xFD:
		return emptyPrefix, uint8(val)
	case val <= 0xFFFF:
		return 0xFD, uint16(val)
	case val <= 0xFFFFFFFF:
		return 0xFE, uint32(val)
	default:
		return 0xFF, uint64(val)
	}
}

// compactSizeSize returns the size in bytes of the encoded unsigned integer
// as specified by the prefix for the CompactSize. The method is used to
// deduce the number of bytes to read after determining the prefix.
func compactSizeSize(prefix byte, pver uint32) uint32 {
	switch prefix {
	case 0xFD:
		return 3
	case 0xFE:
		return 5
	case 0xFF:
		return 9
	default:
		return 1
	}
}
