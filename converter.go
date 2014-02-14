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
