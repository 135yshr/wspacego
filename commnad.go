package wspacego

import ()

type Command struct {
	cmd    string
	subcmd string
	param  int
}

func (c *Command) String() string {
	return "cmd subcmd 1"
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
