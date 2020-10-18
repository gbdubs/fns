package cli

import (
	"flag"
	"github.com/gbdubs/fns/cmd"
	"github.com/gbdubs/fns/util"
)

const pushFunctionCommandName string = "push_fn"

type pushFunctionCliCommand struct {
	dryrun      bool
	interactive bool
}

func (c *pushFunctionCliCommand) Parse(args []string, ctx *CliContext) cmd.Command {
	fs := flag.NewFlagSet(pushFunctionCommandName, flag.ExitOnError)
	dryrun := addDryrunFlag(fs)
	interactive := addInteractiveFlag(fs)
	where := addWhereFlag(fs, ctx)
	project := addProjectRootFlag(fs, ctx)
	fn := addFnRootFlag(fs, ctx)

	fs.Parse(args)

	c.dryrun = *dryrun
	c.interactive = *interactive
	projectRoot := getProjectRootFromFlags(where, project)
	functionRoot := getFunctionRootFromFlags(where, project, fn)
	return &cmd.PushFunction{
		ProjectRoot:  projectRoot,
		FunctionRoot: functionRoot,
	}
}

func (c *pushFunctionCliCommand) Validate(errs *util.ErrList) {
	validateInteractiveAndDryrunNotBothSet(
		errs, c.interactive, c.dryrun)
}

func (c *pushFunctionCliCommand) IsInteractive() bool {
	return c.interactive
}

func (c *pushFunctionCliCommand) IsDryrun() bool {
	return c.dryrun
}
