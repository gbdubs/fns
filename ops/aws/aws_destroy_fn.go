package ops

type AwsDestroyFunctionOp struct {
	Region       string
	FunctionName string
}

func (o *AwsDestroyFunctionOp) Validate() error {
	return nil
}

func (o *AwsDestroyFunctionOp) Execute() error {
	return nil
}
