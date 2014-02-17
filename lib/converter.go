package lib

import (
	"bytes"
	"fmt"
)

func createFunction(b byte) (func([]byte) (*Command, int, error), error) {
	switch b {
	case Space:
		return stackManipulation, nil
	case Lf:
		return flowControl, nil
	case Tab:
		return generateFunctions, nil
	}
	return nil, fmt.Errorf("not defined")
}

func stackManipulation(data []byte) (*Command, int, error) {
	if data[0] == Space {
		buf, seek := readEndLf(data[1:])
		num := parseInt(buf)
		return newSubCommandWithParam("stack", "push", num), seek + 1, nil
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
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		buf, seek := readEndLf(data[2:])
		num := parseInt(buf)
		return newSubCommandWithParam("stack", "ncopy", num), seek + 2, nil
	case bytes.Compare(cmd, []byte{Tab, Lf}) == 0:
		buf, seek := readEndLf(data[2:])
		num := parseInt(buf)
		return newSubCommandWithParam("stack", "move", num), seek + 2, nil
	default:
		return nil, 0, fmt.Errorf("not defined command [%s]", "mani")
	}
	return newSubCommand(word, subcmd), len(cmd), nil
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
		return newCommand("return"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Lf, Lf}) == 0:
		return newCommand("exit"), len(cmd), nil
	default:
		return nil, 0, fmt.Errorf("not defined command [%s]", "flow")
	}

	buf, seek := readEndLf(data[len(cmd):])
	subcmd := string(parseZeroOne(buf))

	return newSubCommand(word, subcmd), len(cmd) + seek, nil
}

func generateFunctions(data []byte) (*Command, int, error) {
	var cmd *Command
	var seek int
	var err error
	switch data[0] {
	case Space:
		cmd, seek, err = arithmetic(data[1:])
	case Tab:
		cmd, seek, err = heapAccess(data[1:])
	case Lf:
		cmd, seek, err = i_o(data[1:])
	default:
		return nil, 0, fmt.Errorf("not defined command [%s]", "subimp")
	}
	return cmd, seek + 1, err
}

func arithmetic(data []byte) (*Command, int, error) {
	cmd := data[0:2]
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		return newCommand("add"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		return newCommand("sub"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Space, Lf}) == 0:
		return newCommand("mul"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		return newCommand("div"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		return newCommand("mod"), len(cmd), nil
	}
	return nil, 0, fmt.Errorf("not defined command [%s]", "arithmetic")
}

func heapAccess(data []byte) (*Command, int, error) {
	const cmd = "heap"
	switch data[0] {
	case Space:
		return newSubCommand(cmd, "push"), 1, nil
	case Tab:
		return newSubCommand(cmd, "pop"), 1, nil
	}
	return nil, 0, fmt.Errorf("not defined command [%s]", "heap")
}

func i_o(data []byte) (*Command, int, error) {
	cmd := data[0:2]
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		return newCommand("putc"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		return newCommand("putn"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		return newCommand("getc"), len(cmd), nil
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		return newCommand("getn"), len(cmd), nil
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
