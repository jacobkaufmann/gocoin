package script

import (
	"bytes"

	"github.com/jacobkaufmann/gocoin/pkg/crypto"
)

// An opCode represents an opcode in the Bitcoin scripting language
type opCode int

const (
	// OpDup duplicates the top item on the stack.
	OpDup opCode = iota

	// OpHash160 performs two hashes using SHA256 followed by RIPEMD-160.
	OpHash160

	// OpEqual consumes the top two items from the stack and determines if they
	// are equal or not.
	OpEqual

	// OpVerify consumes the topmost item on the stack, and if that value is zero,
	// it terminates in failure.
	OpVerify

	// OpEqualVerify runs OpEqual and then OpVerify in sequence.
	OpEqualVerify
)

// Execute the operations specified by the opcode.
func (op opCode) executeOp(s *Stack) error {
	switch op {
	case OpDup:
		// Attempt to pop top item from the stack
		data, err := s.Pop()
		if err != nil {
			return err
		}

		// Duplicate the popped item and push both onto the stack
		dup := make([]byte, len(data))
		copy(dup, data)
		s.Push(data)
		s.Push(dup)
	case OpHash160:
		// Attempt to pop top item from stack
		data, err := s.Pop()
		if err != nil {
			return err
		}

		// Perform SHA256 hash followed by RIPEMD-160 and push result
		// onto the stack
		hash := crypto.Hash160(data)
		s.Push(hash)
	case OpEqual:
		// Attempt to pop top two items from the stack
		d0, err := s.Pop()
		if err != nil {
			return err
		}
		d1, err := s.Pop()
		if err != nil {
			return err
		}

		// If the items are not equal, push a 0 (false) onto the stack.
		// Otherwise, push a 1 (true) onto the stack
		if bytes.Equal(d0, d1) == false {
			s.Push([]byte{0})
		} else {
			s.Push([]byte{1})
		}
	case OpVerify:
		// Attempt to remove the top item from the stack
		data, err := s.Pop()
		if err != nil {
			return err
		}
		// If the item equals false, terminate the script in error
		if bytes.Equal(data, []byte{0}) {
			return err
		}
	case OpEqualVerify:
		err := OpEqual.executeOp(s)
		if err != nil {
			return err
		}
		err = OpVerify.executeOp(s)
		if err != nil {
			return err
		}
	}

	return nil
}
