package cli

import (
	"flag"
	"github.com/gbdubs/fns/cmd"
	"github.com/gbdubs/fns/util"
)

const createProjectCommandName = "create_project"

type createProjectCliCommand struct {
	dryrun      bool
	interactive bool
}

func (c *createProjectCliCommand) Parse(args []string, ctx *CliContext) cmd.Command {
	fs := flag.NewFlagSet(createProjectCommandName, flag.ExitOnError)
	dryrun := addDryrunFlag(fs)
	interactive := addInteractiveFlag(fs)
	project := addNewProjectFlag(fs, ctx)
	where := addWhereFlag(fs, ctx)

	fs.Parse(args)

	c.dryrun = *dryrun
	c.interactive = *interactive
	return &cmd.CreateProject{
		ProjectName: *project,
		ProjectRoot: getProjectRootFromFlags(where, project),
	}
}

func (c *createProjectCliCommand) Validate(errs *util.ErrList) {
	validateInteractiveAndDryrunNotBothSet(
		errs, c.interactive, c.dryrun)
}

func (c *createProjectCliCommand) IsInteractive() bool {
	return c.interactive
}

func (c *createProjectCliCommand) IsDryrun() bool {
	return c.dryrun
}
