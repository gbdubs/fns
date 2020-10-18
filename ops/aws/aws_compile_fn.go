package ops

import (
	"context"
	"errors"
	"fmt"
	"github.com/gbdubs/fns/cfg"
	"os/exec"
	"strings"
	"time"
)

type AwsCompileFunctionOp struct {
	FunctionRoot string
}

func (c *AwsCompileFunctionOp) Validate() error {
	if !cfg.LoadFunctionConfig(c.FunctionRoot).Exists() {
		return errors.New("Function doesn't exist!")
	}
	return nil
}

func (c *AwsCompileFunctionOp) Execute() error {
	buildInputPath := c.FunctionRoot
	buildOutputPath := c.FunctionRoot + ".fndeploy.built"
	zippedOutputPath := ".fndeploy.zip"
	buildArgs := strings.Split(fmt.Sprintf("GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o %s %s", buildInputPath, buildOutputPath), " ")
	ctx, cancelBuild := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelBuild()
	_, err := exec.CommandContext(ctx, "env", buildArgs...).CombinedOutput()
	if err != nil {
		return err
	}

	zipArgs := strings.Split(fmt.Sprintf("-j %s %s", zippedOutputPath, buildOutputPath), " ")
	ctx, cancelZip := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelZip()
	_, err = exec.CommandContext(ctx, "zip", zipArgs...).CombinedOutput()
	return err
}
