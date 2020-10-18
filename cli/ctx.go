package cli

import (
	"github.com/gbdubs/fns/cfg"
	"os"
)

type CliContext struct {
	pwd string
}

func NewCtx() *CliContext {
	ospwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &CliContext{
		pwd: ospwd,
	}
}

// Returns the current project name (if we're working within a project, empty otherwise)
func (c *CliContext) getProjectRoot() string {
	pr := cfg.GetProjectConfigFromPath(c.pwd)
	if cfg.ExistsOnDisk(pr) {
		return cfg.Root(pr)
	}
	return ""
}

// Returns the current function name (if we're working within a function, empty otherwise)
func (c *CliContext) getFunctionRoot() string {
	f := cfg.GetFunctionConfigFromPath(c.pwd)
	if cfg.ExistsOnDisk(f) {
		return cfg.Root(f)
	}
	return ""
}

// Returns the present working directory
func (c *CliContext) getWhere() string {
	return c.pwd
}
