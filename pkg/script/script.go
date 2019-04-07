package script

// A Script is an executable series of Bitcoin scripting language instructions.
type Script struct {
	DataStack    Stack
	Instructions []Instruction
}

// Execute executes the instructions in a script.
func (s Script) Execute() error {
	for _, i := range s.Instructions {
		err := i.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}
