// Variable-length string utilites for serialization and deserialization. The
// variable-length string (VarStr) encoding used by the Bitcoin contains a
// CompactSize followed by the string itself. A separate type is unnecessary
// because we only need the variable-length representation at serialization
// time.

package protocol

import "io"

// writeVarStr encodes s as a variable-length string and writes the value to w.
func writeVarStr(w io.Writer, pver uint32, s string) error {
	strLen := uint64(len(s))
	err := writeCompactSize(w, pver, strLen)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(s))
	if err != nil {
		return err
	}

	return nil
}

// readVarStr reads from r and decodes the variable-length string into a
// string.
func readVarStr(r io.Reader, pver uint32, s *string) error {
	var strLen uint64
	err := readCompactSize(r, pver, &strLen)
	if err != nil {
		return err
	}

	buf := make([]byte, strLen)
	_, err = r.Read(buf)
	if err != nil {
		return err
	}

	*s = string(buf)
	return nil
}
