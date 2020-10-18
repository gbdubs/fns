package files

import (
	"fmt"
	"github.com/gbdubs/fns/util"
	"os"
)

type CreateFolderOp struct {
	FolderPath string
}

func (o *CreateFolderOp) ToDebugString() string {
	return fmt.Sprintf("Create Folder %s", o.FolderPath)
}

func (o *CreateFolderOp) Validate(errs *util.ErrList) {
	_, e := os.Stat(o.FolderPath)
	if !os.IsNotExist(e) {
		errs.Appendf("cannot create directory at %s, it already exists", o.FolderPath)
	}
}

func (o *CreateFolderOp) Execute() error {
	return os.MkdirAll(o.FolderPath, 0755)
}
