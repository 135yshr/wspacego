package wspacego

type Command struct {
	cmd    string
	subcmd string
	param  int
}

func NewCommand(cmd, subcmd string) *Command {
	return nil
}
