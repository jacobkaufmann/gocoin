package script

import "errors"

// A Stack holds a collection of data from DataInstruction.
type Stack struct {
	Items []Data
}

// Data represents the data from a DataInstruction.
type Data []byte

// Push pushes data onto the top of the Stack.
func (s Stack) Push(d Data) {
	s.Items = append(s.Items, d)
}

// Pop attempts to remove the top item from the Stack and
// returns it if Stack is non-empty.
// Otherwise, an error is returned.
func (s Stack) Pop() ([]byte, error) {
	size := len(s.Items)
	if size == 0 {
		return nil, ErrPopFromEmptyStack
	}
	i := s.Items[size-1]
	s.Items = s.Items[:size-1]
	return i, nil
}

// ErrPopFromEmptyStack is the result of an attempt to pop from an empty stack.
var ErrPopFromEmptyStack = errors.New("cannot pop from empty stack")
