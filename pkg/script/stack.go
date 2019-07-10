package script

import "errors"

// A Stack holds a collection of data from DataInstruction.
type Stack struct {
	items [][]byte
}

// NewStack returns an empty stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push pushes data onto the top of the Stack.
func (s Stack) Push(d []byte) {
	s.items = append(s.items, d)
}

// Pop attempts to remove the top item from the Stack and
// returns it if Stack is non-empty.
// Otherwise, an error is returned.
func (s Stack) Pop() ([]byte, error) {
	size := len(s.items)
	if size == 0 {
		return nil, ErrPopFromEmptyStack
	}
	i := s.items[size-1]
	s.items = s.items[:size-1]
	return i, nil
}

// ErrPopFromEmptyStack is the result of an attempt to pop from an empty stack.
var ErrPopFromEmptyStack = errors.New("cannot pop from empty stack")
