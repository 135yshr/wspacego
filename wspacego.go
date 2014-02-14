package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"wspacego/lib"
)

var (
	showHelp bool
)

func main() {
	flag.Parse()
	if showHelp {
		flag.Usage()
		os.Exit(0)
	}
	if flag.NArg() < 2 {
		fmt.Fprintf(os.Stderr, "missing arguments\n")
		flag.Usage()
		os.Exit(-1)
	}

	filename := flag.Arg(1)
	original, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(-1)
	}

	interpreter := lib.NewInterpreter(original)
	mode := flag.Arg(0)
	switch strings.ToLower(mode) {
	case "run":
		interpreter.Run()
	case "disasm":
		interpreter.PrintCode()
	case "char":
		interpreter.PrintChar()
	default:
		fmt.Fprintf(os.Stderr, "not support subcommand\n")
		flag.Usage()
		os.Exit(-1)
	}
	os.Exit(0)
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: [run|disasm|char] <whitespace file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\trun    run the program\n")
		fmt.Fprintf(os.Stderr, "\tdisasm disassemble the program\n")
		fmt.Fprintf(os.Stderr, "\tchar   convert the program (space -> S, Tab -> T)\n")
		flag.PrintDefaults()
	}
	flag.BoolVar(&showHelp, "h", false, "display this help and exit")
}
