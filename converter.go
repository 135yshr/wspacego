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
