package wspacego

type Interpreter struct {
	path string
}

func NewInterpreter(path string) *Interpreter {
	return &Interpreter{path}
}
