package cli

import (
	"flag"
	"strings"
)

const dryrunFlagName string = "dryrun"
const dryrunFlagHelpText string = "whether or not this command should actually execute, or just print out a summary of what it would do."

func addDryrunFlag(fs *flag.FlagSet) *bool {
	return fs.Bool(dryrunFlagName, false, dryrunFlagHelpText)
}

const interactiveFlagName string = "interactive"
const interactiveFlagHelpText string = "whether this command should print a warning/confirmation before executing the scheduled operations."

func addInteractiveFlag(fs *flag.FlagSet) *bool {
	return fs.Bool(interactiveFlagName, false, interactiveFlagHelpText)
}

const newProjectFlagName string = "project"
const newProjectFlagHelpText string = "the name of the project to create."

func addNewProjectFlag(fs *flag.FlagSet, cc *CliContext) *string {
	return fs.String(newProjectFlagName, "", newProjectFlagHelpText)
}

const projectRootFlagName string = "project"
const projectRootFlagHelpText string = "the project to operate on, defaults to the project found along the path the command was invoked on."

func addProjectRootFlag(fs *flag.FlagSet, cc *CliContext) *string {
	return fs.String(projectRootFlagName, cc.getProjectRoot(), projectRootFlagHelpText)
}

const whereFlagName string = "where"
const whereFlagHelpText string = "where to simulate this command is running from, defaults to the present working directory"

func addWhereFlag(fs *flag.FlagSet, ctx *CliContext) *string {
	return fs.String(whereFlagName, ctx.getWhere(), whereFlagHelpText)
}

const newFnFlagName string = "fn"
const newFnFlagHelpText string = "the name of the function to create."

func addNewFnFlag(fs *flag.FlagSet) *string {
	return fs.String(newFnFlagName, "", newFnFlagHelpText)
}

const fnRootFlagName string = "fn"
const fnRootFlagHelpText string = "the name of the function to operate over."

func addFnRootFlag(fs *flag.FlagSet, ctx *CliContext) *string {
	return fs.String(fnRootFlagName, ctx.getFunctionRoot(), fnRootFlagHelpText)
}

const fnTypeFlagName string = "type"
const fnTypeFlagHelpText string = "the type of the function to create."

func addFnTypeFlag(fs *flag.FlagSet) *string {
	return fs.String(fnTypeFlagName, "", fnTypeFlagHelpText)
}

func fixWhereFlag(s *string) string {
	if !strings.HasSuffix(*s, "/") {
		return *s + "/"
	}
	return *s
}

func getProjectRootFromFlags(w *string, p *string) string {
	if *p == "" {
		return ""
	}
	if strings.HasPrefix(*p, *w) || strings.HasPrefix(*w, *p) {
		return *p
	}
	return fixWhereFlag(w) + *p
}

func getFunctionRootFromFlags(w *string, p *string, f *string) string {
	if *f == "" {
		return ""
	}
	if strings.HasPrefix(*f, *w) {
		return *f
	}
	return getProjectRootFromFlags(w, p) + "/" + *f
}
