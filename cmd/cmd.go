package cmd

import (
	"github.com/gbdubs/fns/ops"
	"github.com/gbdubs/fns/util"
)

type Command interface {
	Validate(*util.ErrList)
	ToOpTree() *ops.OpTree
}
