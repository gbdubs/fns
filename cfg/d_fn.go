package cfg

import (
	"github.com/iancoleman/strcase"
)

type DerivedFunctionConfig struct {
	cd                       configData
	FunctionConfig           FunctionConfig
	DerivedProjectConfig     DerivedProjectConfig
	DerivedAwsFunctionConfig DerivedAwsFunctionConfig
	CodeHash                 string
}

func (c *DerivedFunctionConfig) cfg() *configData {
	return &c.cd
}

func (c *DerivedFunctionConfig) fileName() string {
	return ".deployed.function.json"
}

func (f *DerivedFunctionConfig) SimplePackageName() string {
	return f.FunctionConfig.Name
}

func (f *DerivedFunctionConfig) FullPackageName() string {
	return Root(&f.FunctionConfig)
}

func (f *DerivedFunctionConfig) MessagePrefix() string {
	return strcase.ToCamel(f.FunctionConfig.Name)
}

func (f *DerivedFunctionConfig) HasPlatformConfig() bool {
	return IsEmpty(&f.DerivedAwsFunctionConfig, &DerivedAwsFunctionConfig{})
}

func LoadDeployedDerivedFunctionConfig(functionRoot string) *DerivedFunctionConfig {
	c := &DerivedFunctionConfig{}
	LoadDerived(c, functionRoot)
	return c
}
