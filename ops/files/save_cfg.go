package files

import (
	"fmt"
	"github.com/gbdubs/fns/cfg"
	"github.com/gbdubs/fns/util"
)

type SaveConfigOp struct {
	Config cfg.Config
}

func (o *SaveConfigOp) ToDebugString() string {
	return fmt.Sprintf("Save Config to %s", cfg.Path(o.Config))
}

func (o *SaveConfigOp) Validate(errs *util.ErrList) {}

func (o *SaveConfigOp) Execute() error {
	cfg.Save(o.Config)
	return nil
}
