package cfg

import (
	"github.com/gbdubs/fns/util"
	"strings"
)

type FunctionConfig struct {
	cd   configData
	Name string
	Type string
}

func (f *FunctionConfig) cfg() *configData {
	return &f.cd
}

func (f *FunctionConfig) fileName() string {
	return "function.json"
}

func (f *FunctionConfig) ToDerived() *DerivedFunctionConfig {
	p := GetProjectConfigFromPath(Root(f)).ToDerived()
	return &DerivedFunctionConfig{
		cd:                   *f.cd.ToDerived(),
		FunctionConfig:       *f,
		DerivedProjectConfig: *p,
		CodeHash:             util.ComputeFileHash(Path(f)),
	}
}

func LoadFunctionConfig(functionRoot string) *FunctionConfig {
	f := &FunctionConfig{}
	Load(f, functionRoot)
	return f
}

func GetFunctionConfigFromPath(path string) *FunctionConfig {
	if len(path) <= 0 {
		return &FunctionConfig{}
	}
	f := LoadFunctionConfig(path)
	if ExistsOnDisk(f) {
		return f
	}
	return GetFunctionConfigFromPath(path[:(strings.LastIndex(path, "/"))])
}
