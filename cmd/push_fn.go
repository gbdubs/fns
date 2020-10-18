package cmd

import (
	"github.com/gbdubs/fns/cfg"
	"github.com/gbdubs/fns/ops"
	"github.com/gbdubs/fns/util"
)

type PushFunction struct {
	ProjectRoot  string
	FunctionRoot string
	NewPlatforms map[string]bool
}

func (p *PushFunction) Validate(errs *util.ErrList) {
	validateProjectIsSpecifiedAndExists(errs, p.ProjectRoot)
	validateFunctionIsSpecifiedAndExists(errs, p.FunctionRoot)
	if len(p.NewPlatforms) == 0 {
		errs.Appendf("No platform could be deterimed. Either specify a platform configuration within the project or function directories, or use the --platform flag to push this function to a new platform.")
	}
}

func (p *PushFunction) ToOpTree() *ops.OpTree {
	local := cfg.LoadFunctionConfig(p.FunctionRoot).ToDerived()
	deployed := cfg.LoadDeployedDerivedFunctionConfig(p.FunctionRoot)
	platformPushes := []*ops.OpTree{}
	platformPushes = append(platformPushes, PushFnToAws(deployed, local))
	return ops.InParallel(platformPushes...)
}

func PushFnToAws(deployed *cfg.DerivedFunctionConfig, local *cfg.DerivedFunctionConfig) *ops.OpTree {
	return ops.DoNothing()
}
