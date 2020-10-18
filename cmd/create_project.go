package cmd

import (
	"github.com/gbdubs/fns/cfg"
	"github.com/gbdubs/fns/ops"
	"github.com/gbdubs/fns/ops/files"
	"github.com/gbdubs/fns/util"
)

type CreateProject struct {
	ProjectName string
	ProjectRoot string
}

func (c *CreateProject) Validate(errs *util.ErrList) {
	if c.ProjectName == "" {
		errs.Appendf("flag --project must be specified")
	}
	if cfg.ExistsOnDisk(cfg.LoadProjectConfig(c.ProjectRoot)) {
		errs.Appendf("project by this name (%s) already exists.", c.ProjectRoot)
	}
	if cfg.IsEmpty(cfg.GetProjectConfigFromPath(c.ProjectRoot), &cfg.ProjectConfig{}) {
		errs.Appendf("cannot create project within another project.")
	}
}

func (c *CreateProject) ToOpTree() *ops.OpTree {
	projectCfg := cfg.LoadProjectConfig(c.ProjectRoot)
	projectCfg.Name = c.ProjectName
	return ops.DoSequentially(
		&files.CreateFolderOp{
			FolderPath: c.ProjectRoot,
		},
		&files.SaveConfigOp{
			Config: projectCfg,
		},
	)
}
