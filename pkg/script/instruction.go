package script

// An instruction is an executable instruction specified by the
// Bitcoin scripting language.
// Instructions may be data instructions or opcodes.
type instruction interface {
	// Execute the Instruction
	Execute() error
}

// A dataInstruction is a Bitcoin scripting language data instruction.
type dataInstruction struct {
	data  []byte
	stack *Stack
}

// Execute pushes the dataInstruction onto its associated Stack.
func (i *dataInstruction) Execute() error {
	i.stack.Push(i.data)
	return nil
}

// An opCodeInstruction is a Bitcoin scripting language opcode instruction.
type opCodeInstruction struct {
	op    opCode
	stack *Stack
}

// Execute executes the operation in the opCodeInstruction.
func (i *opCodeInstruction) Execute() error {
	return i.op.executeOp(i.stack)
}
