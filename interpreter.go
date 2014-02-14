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
	origin   []byte
	source   []byte
	commands *CommandList
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

func (inter *Interpreter) parseCommands() error {
	inp := *inter
	data := inp.source
	inp.commands = NewCommandList()
	for pos := 0; pos < len(data); pos++ {

		fn, err := CreateFunction(data[pos])
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

func readEndLf(data []byte) ([]byte, int) {
	var ret []byte
	for _, b := range data {
		if b == Lf {
			break
		}
		ret = append(ret, b)
	}
	return ret, len(ret) + 1
}

func parseInt(data []byte) int {
	var ret int
	for _, b := range data {
		ret = ret << 1
		if b == Tab {
			ret += 1
		}
	}
	return ret
}

func parseZeroOne(data []byte) []byte {
	ret := make([]byte, len(data))
	for n, b := range data {
		switch b {
		case Space:
			ret[n] = '0'
		case Tab:
			ret[n] = '1'
		default:
			ret[n] = '-'
		}
	}
	return ret
}
