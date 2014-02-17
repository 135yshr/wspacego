package lib

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func (inter *Interpreter) PrintChar() {
	bys, err := inter.toChar()
	if err != nil {
		panic(err)
	}
	fmt.Print(string(bys))
}

func (inter *Interpreter) PrintCode() {
	err := inter.toCode()
	if err != nil {
		panic(err)
	}
	max := inter.commands.Len()
	for n := 0; n < max; n++ {
		fmt.Println(inter.commands.Get(n + 1))
	}
}

func (inter *Interpreter) Run() {
	err := inter.toCode()
	if err != nil {
		panic(err)
	}

	call_stack := newStack()
	stack := newStack()
	heap := newHeap()

	max := inter.commands.Len()
	for p := 1; p <= max; p++ {
		cmd := inter.commands.Get(p)
		switch cmd.cmd {
		case "stack":
			switch cmd.subcmd {
			case "push":
				stack.Push(cmd.param)
			case "copy":
				b := stack.Pop()
				stack.Push(b)
				stack.Push(b)
			case "ncopy":
			case "swap":
				stack.Swap()
			case "remove":
				stack.Pop()
			}
		case "heap":
			switch cmd.subcmd {
			case "push":
				v := stack.Pop()
				k := stack.Pop()
				heap.Push(k, v)
			case "pop":
				k := stack.Pop()
				stack.Push(heap.Pop(k))
			}
		case "putc":
			fmt.Print(string(stack.Pop()))
		case "putn":
			fmt.Print(stack.Pop())
		case "getc":
			line, err := readStdin()
			if err != nil {
				fmt.Println(err)
				return
			}

			bl := []byte(line)
			for pos := len(bl); 0 < pos; pos-- {
				k := stack.Pop()
				heap.Push(k, int(bl[pos]))
			}
		case "getn":
			line, err := readStdin()
			if err != nil {
				fmt.Println(err)
				return
			}

			num, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
				return
			}

			k := stack.Pop()
			heap.Push(k, num)
		case "return":
			p = call_stack.Pop()
		case "exit":
			fmt.Println("exit program")
			break
		// case "label":
		case "call":
			call_stack.Push(p)
			p, err = inter.commands.Search(newSubCommand("label", cmd.subcmd))
			if err != nil {
				panic(err)
			}
		case "goto":
			p, err = inter.commands.Search(newSubCommand("label", cmd.subcmd))
			if err != nil {
				panic(err)
			}
		case "if stack==0 then goto":
			if stack.Pop() == 0 {
				p, err = inter.commands.Search(newSubCommand("label", cmd.subcmd))
				if err != nil {
					panic(err)
				}
			}
		case "if stack!=0 then goto":
			if stack.Pop() != 0 {
				p, err = inter.commands.Search(newSubCommand("label", cmd.subcmd))
				if err != nil {
					panic(err)
				}
			}
		case "add":
			two, one := stack.Pop(), stack.Pop()
			stack.Push(one + two)
		case "sub":
			two, one := stack.Pop(), stack.Pop()
			stack.Push(one - two)
		case "mul":
			two, one := stack.Pop(), stack.Pop()
			stack.Push(one * two)
		case "div":
			two, one := stack.Pop(), stack.Pop()
			stack.Push(one / two)
		case "mod":
			two, one := stack.Pop(), stack.Pop()
			stack.Push(one % two)
		}
	}
}

func (inter *Interpreter) toChar() ([]byte, error) {
	var ret []byte
	for _, b := range inter.origin {
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

func (inter *Interpreter) toCode() error {
	inter.filter()
	return inter.parseCommands()
}

func (inter *Interpreter) filter() {
	for _, b := range inter.origin {
		if bytes.IndexByte([]byte{Space, Tab, Lf}, b) >= 0 {
			inter.source = append(inter.source, b)
		}
	}
}

func (inter *Interpreter) parseCommands() error {
	data := inter.source
	max := len(data)
	inter.commands = newCommandList()
	for pos := 0; pos < max; {
		fn, err := createFunction(data[pos])
		pos += 1
		command, seek, err := fn(data[pos:])
		if err != nil {
			return err
		}

		pos += seek
		inter.commands.Add(command)
	}
	return nil
}

func readStdin() (string, error) {
	rd := bufio.NewReader(os.Stdin)
	line, err := rd.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.Trim(line, "\r\n"), nil
}
