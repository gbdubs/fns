package cmd

import (
	"github.com/gbdubs/fns/cfg"
	"github.com/gbdubs/fns/ops"
	"github.com/gbdubs/fns/ops/files"
	"github.com/gbdubs/fns/util"
)

type CreateFunction struct {
	ProjectRoot  string
	FunctionName string
	FunctionType string
}

func (c *CreateFunction) Validate(errs *util.ErrList) {
	validateProjectIsSpecifiedAndExists(errs, c.ProjectRoot)
	validateFlagIsSpecified(errs, "fn", c.FunctionName)
	validateFlagIsSpecified(errs, "type", c.FunctionType)
	fnRootDir := c.ProjectRoot + "/" + c.FunctionName
	if cfg.ExistsOnDisk(cfg.LoadFunctionConfig(fnRootDir)) {
		errs.Appendf("cannot create a new function at (%s), one already exists!", fnRootDir)
	}
	if c.FunctionType != "UNAUTH" {
		errs.Appendf("function --type=%s is not supported", c.FunctionType)
	}
}

func (c *CreateFunction) ToOpTree() *ops.OpTree {
	fnRootDir := c.ProjectRoot + "/" + c.FunctionName
	fnCfg := cfg.LoadFunctionConfig(fnRootDir)
	fnCfg.Name = c.FunctionName
	fnCfg.Type = c.FunctionType
	derived := fnCfg.ToDerived()

	createFolderOp := &files.CreateFolderOp{
		FolderPath: fnRootDir,
	}
	createFnCodeOp := &files.CreateFileOp{
		FilePath: fnRootDir + "/function.go",
		Contents: util.RenderTemplate("new_fn.go.tmpl", derived),
	}
	createFnConfigOp := &files.SaveConfigOp{
		Config: fnCfg,
	}
	return ops.Do(createFolderOp).ThenDoInParallel(
		createFnCodeOp,
		createFnConfigOp)
}
