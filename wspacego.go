package main

import (
	"flag"
	"fmt"
	"os"
	"wspacego/lib"
)

func main() {
	flag.Parse()
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
		return
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
	}
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: [run|text|char] <whitespace file>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\trun    run the program\n")
		fmt.Fprintf(os.Stderr, "\tdisasm disassemble the program\n")
		fmt.Fprintf(os.Stderr, "\tchar   convert the program (space -> S, Tab -> T)\n")
		PrintDefaults()
	}
	flag.BoolVar(&showHelp, "h", false, "display this help and exit")
}
