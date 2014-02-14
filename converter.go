package wspacego

import (
	"bytes"
	"fmt"
)

type Converter struct {
}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) CreateFunction(b byte) (func([]byte) (*Command, int, error), error) {
	switch b {
	case Space:
		return c.stackManipulation, nil
	case Lf:
		return c.flowControl, nil
		// case Tab:
		// 	return generateSubImpfunc, nil
	}
	return nil, fmt.Errorf("not defined")
}

func (c *Converter) stackManipulation(data []byte) (*Command, int, error) {
	if data[0] == Space {
		buf, seek := readEndLf(data[1:])
		num := parseInt(buf)
		return NewSubCommandWithParam("stack", "push", num), seek + 1, nil
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
	return NewSubCommand(word, subcmd), len(data), nil
}

func (c *Converter) flowControl(data []byte) (*Command, int, error) {
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
		word = "if stack!=0 then goto"
	case bytes.Compare(cmd, []byte{Tab, Lf}) == 0:
		return NewCommand("return"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Lf, Lf}) == 0:
		return NewCommand("exit"), len(cmd), nil
	default:
		return nil, 0, fmt.Errorf("not defined command [%s]", "flow")
	}

	buf, seek := readEndLf(data[len(cmd):])
	subcmd := string(parseZeroOne(buf))

	return NewSubCommand(word, subcmd), len(cmd) + seek, nil
}

// func (c *Converter) generateSubImpfunc(data []byte) (*Command, int, error) {
// 	switch data[0] {
// 	case Space:
// 		return arithmetic(data[1:])
// 	case Tab:
// 		return heapAccess(data[1:])
// 	case Lf:
// 		return i_o(data[1:])
// 	}
// 	return nil, 0, fmt.Errorf("not defined command [%s]", "subimp")
// }

func (c *Converter) arithmetic(data []byte) (*Command, int, error) {
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

func (c *Converter) heapAccess(data []byte) (*Command, int, error) {
	const cmd = "heap"
	switch data[0] {
	case Space:
		return NewSubCommand(cmd, "push"), 1, nil
	case Tab:
		return NewSubCommand(cmd, "pop"), 1, nil
	}
	return nil, 0, fmt.Errorf("not defined command [%s]", "heap")
}

func (c *Converter) i_o(data []byte) (*Command, int, error) {
	cmd := data[0:2]
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		return NewCommand("putc"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		return NewCommand("putn"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		return NewCommand("getc"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		return NewCommand("getn"), len(cmd), nil
	}
	return nil, 0, fmt.Errorf("not defined command [%s]", "io")
}
