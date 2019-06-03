package protocol

// A VarStr is a variable length string.
type VarStr string

// Size returns the length of the underlying string as a CompactSize.
func (vstr VarStr) Size() CompactSize {
	return CompactSize(len(vstr))
}

func (vstr VarStr) String() string {
	return string(vstr)
}
