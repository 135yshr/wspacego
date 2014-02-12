package wspacego

import (
	"bytes"
	"fmt"
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

		fn, err := generateImpfunc(data[pos])
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

func generateImpfunc(b byte) (func([]byte) (*Command, int, error), error) {
	switch b {
	case Space:
		return stackManipulation, nil
	case Lf:
		return flowControl, nil
	case Tab:
		return generateSubImpfunc, nil
	}
	return nil, fmt.Errorf("not defined")
}

func stackManipulation(data []byte) (*Command, int, error) {
	if data[0] == Space {
		buf, seek := readEndLf(data[1:])
		num := parseInt(buf)
		return NewSubCommandWithParam("stack", "push", num), seek, nil
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
		return nil, 0, fmt.Errorf("not defined command [%s]", "mani")
	}
	return NewSubCommand(word, subcmd), len(cmd) - 1, nil
}

func flowControl(data []byte) (*Command, int, error) {
	cmd := data[0:2]

	var word string
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		word = "label"
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		word = "call"
	case bytes.Compare(cmd, []byte{Space, Lf}) == 0:
		word = "goto"
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		word = "if stack==0 then goto"
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		word = "if stack<0 then goto"
	case bytes.Compare(cmd, []byte{Tab, Lf}) == 0:
		return NewCommand("return"), len(cmd) - 1, nil
	case bytes.Compare(cmd, []byte{Lf, Lf}) == 0:
		return NewCommand("exit"), len(cmd) - 1, nil
	default:
		return nil, 0, fmt.Errorf("not defined command [%s]", "flow")
	}

	buf, seek := readEndLf(data[len(cmd):])
	subcmd := string(parseZeroOne(buf))

	return NewSubCommand(word, subcmd), len(cmd) + seek - 1, nil
}

func generateSubImpfunc(data []byte) (*Command, int, error) {
	switch data[0] {
	case Space:
		return arithmetic(data[1:])
	case Tab:
		return heapAccess(data[1:])
	case Lf:
		return i_o(data[1:])
	}
	return nil, 0, fmt.Errorf("not defined command [%s]", "subimp")
}

func arithmetic(data []byte) (*Command, int, error) {
	cmd := data[0:2]
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		return NewCommand("add"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		return NewCommand("sub"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Space, Lf}) == 0:
		return NewCommand("mul"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		return NewCommand("div"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		return NewCommand("mod"), len(cmd), nil
	}
	return nil, 0, fmt.Errorf("not defined command [%s]", "arithmetic")
}

func heapAccess(data []byte) (*Command, int, error) {
	const cmd = "heap"
	switch data[0] {
	case Space:
		return NewSubCommand(cmd, "push"), 1, nil
	case Tab:
		return NewSubCommand(cmd, "pop"), 1, nil
	}
	return nil, 0, fmt.Errorf("not defined command [%s]", "heap")
}

func i_o(data []byte) (*Command, int, error) {
	cmd := data[0:2]
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		return NewCommand("putc"), len(cmd) - 1, nil
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		return NewCommand("putn"), len(cmd) - 1, nil
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		return NewCommand("getc"), len(cmd) - 1, nil
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		return NewCommand("getn"), len(cmd) - 1, nil
	}
	return nil, 0, fmt.Errorf("not defined command [%s]", "io")
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
