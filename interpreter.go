package wspacego

type Interpreter struct {
	data []byte
}

func NewInterpreter(data []byte) *Interpreter {
	return &Interpreter{data}
}
