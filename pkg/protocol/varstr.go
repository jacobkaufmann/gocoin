package protocol

import "io"

// A VarStr is a variable length string.
type VarStr string

// Serialize serializes vstr and writes to w.
func (vstr *VarStr) Serialize(w io.Writer, pver uint32) error {
	b := []byte(vstr)

	err := vstr.Size().Serialize(w, pver)
	if err != nil {
		return err
	}
	return writeElement(w, b)
}

// Deserialize deserializes data from r into vstr.
func (vstr *VarStr) Deserialize(r io.Reader, pver uint32) error {
	var strLen CompactSize
	err := strLen.Deserialize(r, pver)
	if err != nil {
		return err
	}

	var b [strLen]byte
	err = readElement(r, b)
	if err != nil {
		return err
	}

	vstr = string(b)
	return nil
}

// Size returns the length of the underlying string as a CompactSize.
func (vstr VarStr) Size() CompactSize {
	return CompactSize(len(vstr))
}

func (vstr VarStr) String() string {
	return string(vstr)
}
