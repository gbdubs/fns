package ops

type AwsCreateFunctionOp struct {
	Region       string
	FunctionName string
	DeployZip    string
}

func (o *AwsCreateFunctionOp) Validate() error {
	return nil
}

func (o *AwsCreateFunctionOp) Execute() error {
	return nil
}
