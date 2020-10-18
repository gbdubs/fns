package ops

type AwsUpdateFunctionOp struct {
	Region       string
	FunctionName string
	DeployZip    string
}

func (o *AwsUpdateFunctionOp) Validate() error {
	return nil
}

func (o *AwsUpdateFunctionOp) Execute() error {
	return nil
}
