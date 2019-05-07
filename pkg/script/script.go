package script

// A Script is a sequence of Bitcoin scripting language instructions.
type Script struct {
	dataStack    Stack
	instructions []instruction
}

// Execute executes the instructions in a script.
func (s Script) Execute() error {
	for _, i := range s.instructions {
		err := i.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}
