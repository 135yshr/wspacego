package lib

import (
	"bytes"
)

const (
	Tab   = '\t'
	Space = ' '
	Lf    = '\n'
)

type Interpreter struct {
	origin   []byte
	source   []byte
	commands *CommandList
}

func NewInterpreter(data []byte) *Interpreter {
	return &Interpreter{origin: data}
}

func (inter *Interpreter) ToChar() ([]byte, error) {
	inp := *inter
	var ret []byte
	for _, b := range inp.origin {
		switch b {
		case Space:
			ret = append(ret, 'S')
		case Tab:
			ret = append(ret, 'T')
		case Lf:
			ret = append(ret, Lf)
		}
	}
	return ret, nil
}

func (inter *Interpreter) ToCode() error {
	inter.filter()
	inter.parseCommands()
	return inter.parseCommands()
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

func (inter *Interpreter) parseCommands() error {
	inp := *inter
	data := inp.source
	inp.commands = NewCommandList()
	for pos := 0; pos < len(data); pos++ {

		fn, err := createFunction(data[pos])
		pos += 1
		command, seek, err := fn(data[pos:])
		if err != nil {
			return err
		}

		pos += seek
		inp.commands.Add(command)
	}
	*inter = inp
	return nil
}
