package wspacego

type Command struct {
	cmd    string
	subcmd string
	param  int
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
	return &Command{cmd: cmd}
}
