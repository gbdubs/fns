package ops

type AwsCreateFunctionRoleOp struct {
	Region   string
	RoleName string
}

func (o *AwsCreateFunctionRoleOp) Validate() error {
	return nil
}

func (o *AwsCreateFunctionRoleOp) Execute() error {
	return nil
}
