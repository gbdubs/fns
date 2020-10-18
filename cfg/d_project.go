package cfg

type DerivedProjectConfig struct {
	cd            configData
	ProjectConfig ProjectConfig
}

func (c *DerivedProjectConfig) cfg() *configData {
	return &c.cd
}

func (c *DerivedProjectConfig) fileName() string {
	return ".derived.project.json"
}
