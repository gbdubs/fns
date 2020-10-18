package cfg

import (
	"strings"
)

type ProjectConfig struct {
	cd   configData
	Name string
}

func (c *ProjectConfig) cfg() *configData {
	return &c.cd
}

func (c *ProjectConfig) fileName() string {
	return "project.json"
}

func (p *ProjectConfig) ToDerived() *DerivedProjectConfig {
	return &DerivedProjectConfig{
		cd:            *p.cd.ToDerived(),
		ProjectConfig: *p,
	}
}

func LoadProjectConfig(projectRoot string) *ProjectConfig {
	c := &ProjectConfig{}
	Load(c, projectRoot)
	return c
}

func GetProjectConfigFromPath(path string) *ProjectConfig {
	if len(path) <= 0 {
		return &ProjectConfig{}
	}
	p := LoadProjectConfig(path)
	if ExistsOnDisk(p) {
		return p
	}
	return GetProjectConfigFromPath(path[:(strings.LastIndex(path, "/"))])
}
