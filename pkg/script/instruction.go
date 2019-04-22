package script

// An Instruction is an executable instruction specified by the
// Bitcoin scripting language.
// Instructions may be data instructions or opcodes.
type Instruction interface {
	// Execute the Instruction
	Execute() error
}

// A DataInstruction is a Bitcoin scripting language data instruction.
type DataInstruction struct {
	data  []byte
	stack *Stack
}

// Execute pushes the DataInstruction onto its associated Stack.
func (i *DataInstruction) Execute() error {
	i.stack.Push(i.data)
	return nil
}

// An OpCodeInstruction is a Bitcoin scripting language opcode instruction.
type OpCodeInstruction struct {
	op    opCode
	stack *Stack
}

// Execute executes the operation in the OpCodeInstruction.
func (i *OpCodeInstruction) Execute() error {
	return i.op.executeOp(i.stack)
}
