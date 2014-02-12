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

func (inter *Interpreter) parseCommands() {
	inp := *inter
	data := inp.source
	inp.commands = NewCommandList()
	for pos := 0; pos < len(data); pos++ {
		var seek int
		var command *Command

		switch data[pos] {
		case Space:
			command, seek = stackManipulation(data[pos+1:])
		case Lf:
		case Tab:
		default:
			continue
		}
		pos += (seek + 1)
		inp.commands.Add(command)
	}
	*inter = inp
}

func stackManipulation(data []byte) (*Command, int) {
	if data[0] == Space {
		buf, seek := readEndLf(data[1:])
		num := parseInt(buf)
		return NewSubCommandWithParam("stack", "push", num), seek
	}

	var word, subcmd string
	cmd := data[0:2]
	switch {
	case bytes.Compare(cmd, []byte{Lf, Space}) == 0:
		word = "stack"
		subcmd = "copy"
	case bytes.Compare(cmd, []byte{Lf, Tab}) == 0:
		word = "stack"
		subcmd = "swap"
	case bytes.Compare(cmd, []byte{Lf, Lf}) == 0:
		word = "stack"
		subcmd = "remove"
	default:
		return NewCommand("mani.undefined"), 0
	}
	return NewSubCommand(word, subcmd), len(cmd) - 1
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
