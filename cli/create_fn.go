package cli

import (
	"flag"
	"github.com/gbdubs/fns/cmd"
	"github.com/gbdubs/fns/util"
)

const createFunctionCommandName = "create_fn"

type createFunctionCliCommand struct {
	dryrun      bool
	interactive bool
}

func (c *createFunctionCliCommand) Parse(args []string, ctx *CliContext) cmd.Command {
	fs := flag.NewFlagSet(createFunctionCommandName, flag.ExitOnError)
	dryrun := addDryrunFlag(fs)
	interactive := addInteractiveFlag(fs)
	project := addProjectRootFlag(fs, ctx)
	where := addWhereFlag(fs, ctx)
	fn := addNewFnFlag(fs)
	fnType := addFnTypeFlag(fs)

	fs.Parse(args)

	c.dryrun = *dryrun
	c.interactive = *interactive
	projectRoot := getProjectRootFromFlags(where, project)
	return &cmd.CreateFunction{
		ProjectRoot:  projectRoot,
		FunctionName: *fn,
		FunctionType: *fnType,
	}
}

func (c *createFunctionCliCommand) Validate(errs *util.ErrList) {
	validateInteractiveAndDryrunNotBothSet(
		errs, c.interactive, c.dryrun)
}

func (c *createFunctionCliCommand) IsInteractive() bool {
	return c.interactive
}

func (c *createFunctionCliCommand) IsDryrun() bool {
	return c.dryrun
}
