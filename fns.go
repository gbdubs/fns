package main

import (
	"bufio"
	"fmt"
	"github.com/gbdubs/fns/cli"
	"github.com/gbdubs/fns/util"
	"os"
	"strings"
)

func main() {
	ctx := cli.NewCtx()
	errs := util.NewErrList()
	firstArg := ""
	if len(os.Args) >= 2 {
		firstArg = os.Args[1]
	}
	cliCmd := cli.GetCliCommand(firstArg, errs)
	ifErrsPrintAndQuit(errs)

	cliCmd.Validate(errs)
	ifErrsPrintAndQuit(errs)

	cmd := cliCmd.Parse(os.Args[2:], ctx)
	cmd.Validate(errs)
	ifErrsPrintAndQuit(errs)

	opt := cmd.ToOpTree()
	opt.ValidateAll(errs)
	ifErrsPrintAndQuit(errs)

	if cliCmd.IsDryrun() {
		fmt.Printf("If wetrun, these operations will be executed\n\n%s\n\n", opt.ToDebugString())
		return
	}
	if cliCmd.IsInteractive() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("If confirmed, these operations will be executed\n\n%s\n\n", opt.ToDebugString())
		fmt.Printf("Execute? [Y/N] ")
		resp, _ := reader.ReadString('\n')
		if resp != "Y" {
			fmt.Print("Cancelled - did not execute.")
		}
	}
	runtimeErrors := opt.Execute()
	ifRuntimeErrsPrintAndQuit(runtimeErrors)
}

func ifErrsPrintAndQuit(errs *util.ErrList) {
	if errs.IsEmpty() {
		return
	}
	fmt.Printf("Encountered Validation Errors:\n%v", errs)
	os.Exit(1)
}

func ifRuntimeErrsPrintAndQuit(errs []error) {
	if len(errs) == 0 {
		return
	}
	fmt.Printf("Encountered %v Errors While Validating:\n", len(errs))
	for _, err := range errs {
		fmt.Printf("  * %s\n", strings.ReplaceAll(errToString(err), "\n", "\n    "))
	}
	os.Exit(1)
}

func errToString(e error) string {
	return fmt.Sprintf("%v", e)
}
