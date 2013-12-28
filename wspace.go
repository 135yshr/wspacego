package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	Space = ' '
	Tab   = '\t'
	LF    = '\n'
)

type Command struct {
	name   string
	subcmd string
	param  int
}

type (
	Stack        []int
	Heap         map[int]int
	CommandSlice []Command
)

var (
	showHelp bool
	DBG      bool
)

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	filename := flag.Arg(1)
	original, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		return
	}

	data := Filter(original)
	mode := flag.Arg(0)
	switch strings.ToLower(mode) {
	case "run":
		Run(data)
	case "text":
		PrintAll(data)
	case "char":
		PrintChar(data)
	default:
		flag.Usage()
	}
}

func init() {
	flag.BoolVar(&showHelp, "h", false, "display this help and exit")
	flag.BoolVar(&DBG, "d", false, "display trace log")
}

func Filter(data []byte) []byte {
	var ret []byte
	for _, b := range data {
		if bytes.IndexByte([]byte{Space, Tab, LF}, b) >= 0 {
			ret = append(ret, b)
		}
	}
	return ret
}

func Run(data []byte) {
	heapMem := make(Heap)
	stackMem := new(Stack)
	callStack := new(Stack)
	commands := ConvertCode(data)

	for p := 0; p < len(commands); p++ {
		cmd := commands[p]
		switch cmd.name {
		case "stack":
			switch cmd.subcmd {
			case "push":
				stackMem.Push(cmd.param)
			case "copy":
				b := stackMem.Pop()
				stackMem.Push(b)
				stackMem.Push(b)
			case "swap":
				max := stackMem.Len() - 1
				stackMem.Swap(max-1, max)
			case "remove":
				stackMem.Pop()
			}
		case "heap":
			switch cmd.subcmd {
			case "push":
				heapMem.Push(stackMem)
			case "pop":
				stackMem.Push(heapMem.Pop(stackMem))
			}
		case "putc":
			fmt.Print(string(stackMem.Pop()))
		case "putn":
			fmt.Print(stackMem.Pop())
		case "getc", "getn":
			rd := bufio.NewReader(os.Stdin)
			var num int
			line, err := rd.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			num, err = strconv.Atoi(line[0 : len(line)-1])
			if err != nil {
				fmt.Println(err)
				return
			}

			k := stackMem.Pop()
			heapMem[k] = num
			if DBG {
				fmt.Println(heapMem)
			}
		case "return":
			p = callStack.Pop()
		case "exit":
			fmt.Println("exit program")
			p = len(commands)
		// case "label":
		case "call":
			callStack.Push(p)
			p = commands.Search(Command{"label", cmd.subcmd, 0})
		case "goto":
			p = commands.Search(Command{"label", cmd.subcmd, 0})
		case "if stack==0 then goto":
			if stackMem.Pop() == 0 {
				p = commands.Search(Command{"label", cmd.subcmd, 0})
			}
		case "if stack<0 then goto":
			if stackMem.Pop() < 0 {
				p = commands.Search(Command{"label", cmd.subcmd, 0})
			}
		case "add":
			one, two := stackMem.Pop(), stackMem.Pop()
			stackMem.Push(one + two)
		case "sub":
			one, two := stackMem.Pop(), stackMem.Pop()
			stackMem.Push(one - two)
			if DBG {
				fmt.Println(one, two)
			}
		case "mul":
			one, two := stackMem.Pop(), stackMem.Pop()
			stackMem.Push(one * two)
		case "div":
			one, two := stackMem.Pop(), stackMem.Pop()
			stackMem.Push(one / two)
		case "mod":
			one, two := stackMem.Pop(), stackMem.Pop()
			stackMem.Push(one % two)
		}
	}
}

func PrintChar(data []byte) {
	for _, b := range data {
		switch b {
		case Space:
			fmt.Print(string('S'))
		case Tab:
			fmt.Print(string('T'))
		case LF:
			fmt.Print(string(LF))
		}
	}
}

func PrintAll(data []byte) {
	commands := ConvertCode(data)
	for _, cmd := range commands {
		fmt.Println(cmd.ToString())
	}
}

func ConvertCode(data []byte) CommandSlice {
	var ret []Command
	for pos := 0; pos < len(data); pos++ {

		var seek int
		var command Command

		switch data[pos] {
		case Space:
			command, seek = StackManipulation(data[pos+1:])
		case LF:
			command, seek = FlowControl(data[pos+1:])
		case Tab:
			pos += 1
			switch data[pos] {
			case Space:
				command, seek = Arithmetic(data[pos+1:])
			case Tab:
				command, seek = HeapAccess(data[pos+1:])
			case LF:
				command, seek = I_O(data[pos+1:])
			}
		default:
			continue
		}
		pos += seek + 1
		ret = append(ret, command)
	}
	return ret
}

func StackManipulation(data []byte) (Command, int) {
	if data[0] == Space {
		buf, seek := ReadEndLf(data[1:])
		num := ParseInt(buf)
		return Command{"stack", "push", num}, seek
	}

	var word, subcmd string
	cmd := data[0:2]
	switch {
	case bytes.Compare(cmd, []byte{LF, Space}) == 0:
		word = "stack"
		subcmd = "copy"
	case bytes.Compare(cmd, []byte{LF, Tab}) == 0:
		word = "stack"
		subcmd = "swap"
	case bytes.Compare(cmd, []byte{LF, LF}) == 0:
		word = "stack"
		subcmd = "remove"
	default:
		return Command{"mani.undefined", "", 0}, 0
	}
	return Command{word, subcmd, 0}, len(cmd) - 1
}

func FlowControl(data []byte) (Command, int) {
	cmd := data[0:2]

	var word string
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		word = "label"
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		word = "call"
	case bytes.Compare(cmd, []byte{Space, LF}) == 0:
		word = "goto"
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		word = "if stack==0 then goto"
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		word = "if stack<0 then goto"
	case bytes.Compare(cmd, []byte{Tab, LF}) == 0:
		return Command{"return", "", 0}, len(cmd) - 1
	case bytes.Compare(cmd, []byte{LF, LF}) == 0:
		return Command{"exit", "", 0}, len(cmd) - 1
	default:
		return Command{"flow.undefined", "", 0}, 0
	}

	buf, seek := ReadEndLf(data[len(cmd):])
	subcmd := string(ParseZeroOne(buf))

	return Command{word, subcmd, 0}, len(cmd) + seek - 1
}

func Arithmetic(data []byte) (Command, int) {
	cmd := data[0:2]
	var word string
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		word = "add"
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		word = "sub"
	case bytes.Compare(cmd, []byte{Space, LF}) == 0:
		word = "mul"
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		word = "div"
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		word = "mod"
	default:
		return Command{"arit.undefined", "", 0}, 0
	}
	return Command{word, "", 0}, len(cmd) - 1
}

func HeapAccess(data []byte) (Command, int) {
	switch data[0] {
	case Space:
		return Command{"heap", "push", 0}, 0
	case Tab:
		return Command{"heap", "pop", 0}, 0
	default:
		return Command{"heap.undefined", "", 0}, 0
	}
}
func I_O(data []byte) (Command, int) {
	cmd := data[0:2]
	var word string
	switch {
	case bytes.Compare(cmd, []byte{Space, Space}) == 0:
		word = "putc"
	case bytes.Compare(cmd, []byte{Space, Tab}) == 0:
		word = "putn"
	case bytes.Compare(cmd, []byte{Tab, Space}) == 0:
		word = "getc"
	case bytes.Compare(cmd, []byte{Tab, Tab}) == 0:
		word = "getn"
	default:
		return Command{"io.undefined", "", 0}, 0
	}
	return Command{word, "", 0}, len(cmd) - 1
}

func ReadEndLf(data []byte) ([]byte, int) {
	var ret []byte
	for _, b := range data {
		if b == LF {
			break
		}
		ret = append(ret, b)
	}
	return ret, len(ret) + 1
}

func ParseInt(data []byte) int {
	var ret int
	for _, b := range data {
		ret = ret << 1
		if b == Tab {
			ret += 1
		}
	}
	return ret
}

func ParseZeroOne(data []byte) []byte {
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

func (c *Command) ToString() string {
	return fmt.Sprintf("%s %s %d", c.name, c.subcmd, c.param)
}

func (h Stack) Len() int {
	return len(h)
}

func (h Stack) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Stack) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Stack) Push(x int) {
	*h = append(*h, x)
	if DBG {
		fmt.Println(*h)
	}
}

func (h *Stack) Pop() int {
	old := *h
	n := len(old)
	ret := old[n-1]
	*h = old[0 : n-1]
	if DBG {
		fmt.Println(ret)
		fmt.Println(*h)
	}
	return ret
}

func (cs CommandSlice) Search(cmd Command) int {
	for n, c := range cs {
		if c == cmd {
			return n
		}
	}
	return -1
}

func (h Heap) Push(mem *Stack) {
	v := mem.Pop()
	k := mem.Pop()
	h[k] = v
	if DBG {
		fmt.Println(h)
	}
}

func (h Heap) Pop(mem *Stack) int {
	k := mem.Pop()
	v := h[k]
	delete(h, k)
	if DBG {
		fmt.Println(h)
	}
	return v
}
