package wspacego

type Command struct {
	cmd    string
	subcmd string
	param  int
}

func NewCommand(cmd string) *Command {
	return nil
}

func NewSubCommand(cmd, subcmd string) *Command {
	return &Command{cmd: cmd, subcmd: subcmd}
}
