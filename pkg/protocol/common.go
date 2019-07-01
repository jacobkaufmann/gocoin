package protocol

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

var (
	// littleEndian is a convenience variable for binary.littleEndian.
	littleEndian = binary.LittleEndian

	// bigEndian is a convenience variable for binary.BigEndian.
	bigEndian = binary.BigEndian

	// binaryFreeListMaxItems is the number of buffers to keep in the free
	// list to use for binary serialization and deserialization.
	binaryFreeListMaxItems = 1024
)

// binaryFreeList from github.com/btcsuite/btcd.
//
// binaryFreeList defines a concurrent safe free list of byte slices (up to the
// maximum number defined by the binaryFreeListMaxItems constant) that have a
// cap of 8 (thus it supports up to a uint64).  It is used to provide temporary
// buffers for serializing and deserializing primitive numbers to and from their
// binary encoding in order to greatly reduce the number of allocations
// required.
//
// For convenience, functions are provided for each of the primitive unsigned
// integers that automatically obtain a buffer from the free list, perform the
// necessary binary conversion, read from or write to the given io.Reader or
// io.Writer, and return the buffer to the free list.
type binaryFreeList chan []byte

// Borrow returns a byte slice from the free list with a length of 8.  A new
// buffer is allocated if there are not any available on the free list.
func (l binaryFreeList) Borrow() []byte {
	var buf []byte
	select {
	case buf = <-l:
	default:
		buf = make([]byte, 8)
	}
	return buf[:8]
}

// Return puts the provided byte slice back on the free list.  The buffer MUST
// have been obtained via the Borrow function and therefore have a cap of 8.
func (l binaryFreeList) Return(buf []byte) {
	select {
	case l <- buf:
	default:
		// Let it go to the garbage collector.
	}
}

// Uint8 reads a single byte from the provided reader using a buffer from the
// free list and returns it as a uint8.
func (l binaryFreeList) Uint8(r io.Reader) (uint8, error) {
	buf := l.Borrow()[:1]
	if _, err := io.ReadFull(r, buf); err != nil {
		l.Return(buf)
		return 0, err
	}
	rv := buf[0]
	l.Return(buf)
	return rv, nil
}

// Uint16 reads two bytes from the provided reader using a buffer from the
// free list, converts it to a number using the provided byte order, and returns
// the resulting uint16.
func (l binaryFreeList) Uint16(r io.Reader, byteOrder binary.ByteOrder) (uint16, error) {
	buf := l.Borrow()[:2]
	if _, err := io.ReadFull(r, buf); err != nil {
		l.Return(buf)
		return 0, err
	}
	rv := byteOrder.Uint16(buf)
	l.Return(buf)
	return rv, nil
}

// Uint32 reads four bytes from the provided reader using a buffer from the
// free list, converts it to a number using the provided byte order, and returns
// the resulting uint32.
func (l binaryFreeList) Uint32(r io.Reader, byteOrder binary.ByteOrder) (uint32, error) {
	buf := l.Borrow()[:4]
	if _, err := io.ReadFull(r, buf); err != nil {
		l.Return(buf)
		return 0, err
	}
	rv := byteOrder.Uint32(buf)
	l.Return(buf)
	return rv, nil
}

// Uint64 reads eight bytes from the provided reader using a buffer from the
// free list, converts it to a number using the provided byte order, and returns
// the resulting uint64.
func (l binaryFreeList) Uint64(r io.Reader, byteOrder binary.ByteOrder) (uint64, error) {
	buf := l.Borrow()[:8]
	if _, err := io.ReadFull(r, buf); err != nil {
		l.Return(buf)
		return 0, err
	}
	rv := byteOrder.Uint64(buf)
	l.Return(buf)
	return rv, nil
}

// PutUint8 copies the provided uint8 into a buffer from the free list and
// writes the resulting byte to the given writer.
func (l binaryFreeList) PutUint8(w io.Writer, val uint8) error {
	buf := l.Borrow()[:1]
	buf[0] = val
	_, err := w.Write(buf)
	l.Return(buf)
	return err
}

// PutUint16 serializes the provided uint16 using the given byte order into a
// buffer from the free list and writes the resulting two bytes to the given
// writer.
func (l binaryFreeList) PutUint16(w io.Writer, byteOrder binary.ByteOrder, val uint16) error {
	buf := l.Borrow()[:2]
	byteOrder.PutUint16(buf, val)
	_, err := w.Write(buf)
	l.Return(buf)
	return err
}

// PutUint32 serializes the provided uint32 using the given byte order into a
// buffer from the free list and writes the resulting four bytes to the given
// writer.
func (l binaryFreeList) PutUint32(w io.Writer, byteOrder binary.ByteOrder, val uint32) error {
	buf := l.Borrow()[:4]
	byteOrder.PutUint32(buf, val)
	_, err := w.Write(buf)
	l.Return(buf)
	return err
}

// PutUint64 serializes the provided uint64 using the given byte order into a
// buffer from the free list and writes the resulting eight bytes to the given
// writer.
func (l binaryFreeList) PutUint64(w io.Writer, byteOrder binary.ByteOrder, val uint64) error {
	buf := l.Borrow()[:8]
	byteOrder.PutUint64(buf, val)
	_, err := w.Write(buf)
	l.Return(buf)
	return err
}

// binarySerializer provides a free list of buffers to use for serializing and
// deserializing primitive integer values to and from io.Readers and io.Writers.
var binarySerializer binaryFreeList = make(chan []byte, binaryFreeListMaxItems)

// uint32Time represents a unix timestamp encoded with a uint32.
type uint32Time time.Time

// int64Time represents a unix timestamp encoded with an int64.
type int64Time time.Time

func readElement(r io.Reader, element interface{}) error {
	switch e := element.(type) {
	case *bool:
		rv, err := binarySerializer.Uint8(r)
		if err != nil {
			return err
		}
		if rv == 0x00 {
			*e = false
		} else {
			*e = true
		}
		return nil
	case *uint32:
		rv, err := binarySerializer.Uint32(r, littleEndian)
		if err != nil {
			return err
		}
		*e = rv
		return nil
	case *int32:
		rv, err := binarySerializer.Uint32(r, littleEndian)
		if err != nil {
			return err
		}
		*e = int32(rv)
		return nil
	case *uint64:
		rv, err := binarySerializer.Uint64(r, littleEndian)
		if err != nil {
			return err
		}
		*e = rv
		return nil
	case *int64:
		rv, err := binarySerializer.Uint64(r, littleEndian)
		if err != nil {
			return err
		}
		*e = int64(rv)
		return nil
	case *uint32Time:
		rv, err := binarySerializer.Uint32(r, littleEndian)
		if err != nil {
			return err
		}
		*e = uint32Time(time.Unix(int64(rv), 0))
		return nil
	case *int64Time:
		rv, err := binarySerializer.Uint64(r, littleEndian)
		if err != nil {
			return err
		}
		*e = int64Time(time.Unix(int64(rv), 0))
		return nil
	case *[ChecksumSize]byte:
		_, err := io.ReadFull(r, e[:])
		if err != nil {
			return err
		}
		return nil
	case *[CommandSize]byte:
		_, err := io.ReadFull(r, e[:])
		if err != nil {
			return err
		}
		return nil
	case *[net.IPv6len]byte:
		_, err := io.ReadFull(r, e[:])
		if err != nil {
			return err
		}
		return nil
	case *[HashSize]byte:
		_, err := io.ReadFull(r, e[:])
		if err != nil {
			return err
		}
		return nil
	case *ServiceFlag:
		rv, err := binarySerializer.Uint64(r, littleEndian)
		if err != nil {
			return err
		}
		*e = ServiceFlag(rv)
		return nil
	case *InvType:
		rv, err := binarySerializer.Uint32(r, littleEndian)
		if err != nil {
			return err
		}
		*e = InvType(rv)
		return nil
	case *BitcoinNet:
		rv, err := binarySerializer.Uint32(r, littleEndian)
		if err != nil {
			return err
		}
		*e = BitcoinNet(rv)
		return nil
	case *RejectCode:
		rv, err := binarySerializer.Uint8(r)
		if err != nil {
			return err
		}
		*e = RejectCode(rv)
		return nil
	}

	return binary.Read(r, littleEndian, element)
}

// readElements reads multiple items from r.  It is equivalent to multiple
// calls to readElement.
func readElements(r io.Reader, elements ...interface{}) error {
	for _, element := range elements {
		err := readElement(r, element)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeElement(w io.Writer, element interface{}) error {
	switch e := element.(type) {
	case bool:
		var err error
		if e {
			err = binarySerializer.PutUint8(w, 0x01)
		} else {
			err = binarySerializer.PutUint8(w, 0x00)
		}
		if err != nil {
			return err
		}
		return nil
	case uint32:
		err := binarySerializer.PutUint32(w, littleEndian, e)
		if err != nil {
			return err
		}
		return nil
	case int32:
		err := binarySerializer.PutUint32(w, littleEndian, uint32(e))
		if err != nil {
			return err
		}
		return nil
	case uint64:
		err := binarySerializer.PutUint64(w, littleEndian, e)
		if err != nil {
			return err
		}
		return nil
	case int64:
		err := binarySerializer.PutUint64(w, littleEndian, uint64(e))
		if err != nil {
			return err
		}
		return nil
	case [ChecksumSize]byte:
		_, err := w.Write(e[:])
		if err != nil {
			return err
		}
		return nil
	case [CommandSize]byte:
		_, err := w.Write(e[:])
		if err != nil {
			return err
		}
		return nil
	case [net.IPv6len]byte:
		_, err := w.Write(e[:])
		if err != nil {
			return err
		}
		return nil
	case *[HashSize]byte:
		_, err := w.Write(e[:])
		if err != nil {
			return err
		}
		return nil
	case ServiceFlag:
		err := binarySerializer.PutUint64(w, littleEndian, uint64(e))
		if err != nil {
			return err
		}
		return nil
	case InvType:
		err := binarySerializer.PutUint32(w, littleEndian, uint32(e))
		if err != nil {
			return err
		}
		return nil
	case BitcoinNet:
		err := binarySerializer.PutUint32(w, littleEndian, uint32(e))
		if err != nil {
			return err
		}
		return nil
	case RejectCode:
		err := binarySerializer.PutUint8(w, uint8(e))
		if err != nil {
			return err
		}
		return nil
	}

	return binary.Write(w, littleEndian, element)
}

// writeElements writes multiple items to w.  It is equivalent to multiple
// calls to writeElement.
func writeElements(w io.Writer, elements ...interface{}) error {
	for _, element := range elements {
		err := writeElement(w, element)
		if err != nil {
			return err
		}
	}
	return nil
}

// reverseBytes reverses the elements of a byte slice.
func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
