package cfg

type DerivedAwsFunctionConfig struct {
	cd                configData
	AwsFunctionConfig AwsFunctionConfig
}

func (c *DerivedAwsFunctionConfig) cfg() *configData {
	return &c.cd
}

func (c *DerivedAwsFunctionConfig) fileName() string {
	return ".derived.aws.json"
}
