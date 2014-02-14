package lib

import (
	"fmt"
	"reflect"
)

type Command struct {
	cmd    string
	subcmd string
	param  int
}

type CommandList map[int]*Command

func (c *Command) String() string {
	return fmt.Sprintf("%s %s %d", c.cmd, c.subcmd, c.param)
}

func NewCommand(cmd string) *Command {
	return &Command{cmd: cmd}
}

func NewSubCommand(cmd, subcmd string) *Command {
	return &Command{cmd: cmd, subcmd: subcmd}
}

func NewCommandWithParam(cmd string, param int) *Command {
	return &Command{cmd: cmd, param: param}
}

func NewSubCommandWithParam(cmd, subcmd string, param int) *Command {
	return &Command{cmd, subcmd, param}
}

func NewCommandList() *CommandList {
	return &CommandList{}
}

func (l *CommandList) Add(c *Command) {
	cl := *l
	k := len(cl) + 1
	cl[k] = c
	l = &cl
}

func (l *CommandList) Clear() {
	cl := *l
	for k := range cl {
		delete(cl, k)
	}
	l = &cl
}

func (l *CommandList) Get(n int) *Command {
	cl := *l
	return cl[n]
}

func (l *CommandList) Len() int {
	return len(*l)
}

func (l *CommandList) Search(cmd *Command) (int, error) {
	for k, c := range *l {
		if reflect.DeepEqual(c, cmd) {
			return k, nil
		}
	}
	return -1, fmt.Errorf("not defined")
}
