package cfg

type AwsFunctionConfig struct {
	cd     configData
	Region string
}

func (c *AwsFunctionConfig) cfg() *configData {
	return &c.cd
}

func (c *AwsFunctionConfig) fileName() string {
	return "aws.function.json"
}

func (f *AwsFunctionConfig) ToDerived() *DerivedAwsFunctionConfig {
	return &DerivedAwsFunctionConfig{
		cd:                *f.cd.ToDerived(),
		AwsFunctionConfig: *f,
	}
}

func LoadAwsFunctionConfig(functionRoot string) *AwsFunctionConfig {
	f := &AwsFunctionConfig{}
	Load(f, functionRoot)
	return f
}
