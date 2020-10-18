package files

import (
	"fmt"
	"github.com/gbdubs/fns/util"
	"io/ioutil"
	"os"
)

type CreateFileOp struct {
	FilePath string
	Contents string
}

func (c *CreateFileOp) ToDebugString() string {
	return fmt.Sprintf("Create File %s", c.FilePath)
}

func (c *CreateFileOp) Validate(errs *util.ErrList) {
	_, e := os.Stat(c.FilePath)
	if !os.IsNotExist(e) {
		errs.Appendf("cannot create a new file at %s, it already exists", c.FilePath)
	}
}

func (c *CreateFileOp) Execute() error {
	return ioutil.WriteFile(c.FilePath, []byte(c.Contents), 0777)
}
