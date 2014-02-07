package wspacego

import "fmt"

type Command struct {
	cmd    string
	subcmd string
	param  int
}

type CommandList struct {
}

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
