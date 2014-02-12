package wspacego

import (
	"bytes"
)

const (
	Tab   = '\t'
	Space = ' '
	Lf    = '\n'
)

type Interpreter struct {
	origin []byte
	source []byte
}

func NewInterpreter(data []byte) *Interpreter {
	return &Interpreter{origin: data}
}

func (inter *Interpreter) filter() {
	inp := *inter
	for _, b := range inp.origin {
		if bytes.IndexByte([]byte{Space, Tab, Lf}, b) >= 0 {
			inp.source = append(inp.source, b)
		}
	}
	*inter = inp
}
