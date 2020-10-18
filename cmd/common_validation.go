package cmd

import (
	"github.com/gbdubs/fns/cfg"
	"github.com/gbdubs/fns/util"
)

func validateFlagIsSpecified(errs *util.ErrList, flagName string, flagValue string) {
	if flagValue == "" {
		errs.Appendf("flag --%s is required, but wasn't found.", flagName)
	}
}

func validateProjectIsSpecifiedAndExists(errs *util.ErrList, projectRoot string) {
	if projectRoot == "" {
		errs.Appendf("flag --project is required if you are not operating from within a project directory")
	} else if !cfg.ExistsOnDisk(cfg.LoadProjectConfig(projectRoot)) {
		errs.Appendf("no project named %s currently exists.", projectRoot)
	}
}

func validateFunctionIsSpecifiedAndExists(errs *util.ErrList, functionRoot string) {
	if functionRoot == "" {
		errs.Appendf("flag --fn is required if you are not operating from within a function directory")
	} else if !cfg.ExistsOnDisk(cfg.LoadFunctionConfig(functionRoot)) {
		errs.Appendf("no project named %s currently exists", functionRoot)
	}
}
