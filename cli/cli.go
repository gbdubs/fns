package cli

import (
	"github.com/gbdubs/fns/cmd"
	"github.com/gbdubs/fns/util"
)

type CliCommand interface {
	Parse([]string, *CliContext) cmd.Command
	Validate(*util.ErrList)
	IsDryrun() bool
	IsInteractive() bool
}

func cliCommandMap() map[string]CliCommand {
	return map[string]CliCommand{
		createProjectCommandName:  &createProjectCliCommand{},
		createFunctionCommandName: &createFunctionCliCommand{},
		pushFunctionCommandName:   &pushFunctionCliCommand{},
	}
}

func GetCliCommand(c string, errs *util.ErrList) CliCommand {
	m := cliCommandMap()
	if c == "" {
		errs.Appendf("Every call to fns requires a command. Use the help command to learn more: `fns help`")
	} else if m[c] == nil {
		errs.Appendf("No command found for '%s'. Try 'help' for a list of commands", c)
	}
	return m[c]
}
