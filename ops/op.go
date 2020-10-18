package ops

import (
	"github.com/gbdubs/fns/util"
	"strings"
)

type Op interface {
	ToDebugString() string
	Validate(*util.ErrList)
	Execute() error
}

type OpTree struct {
	hasOp    bool
	op       Op
	children []*OpTree
}

func Do(op Op) *OpTree {
	return &OpTree{
		op:       op,
		hasOp:    true,
		children: []*OpTree{},
	}
}

func DoNothing() *OpTree {
	return &OpTree{
		hasOp:    false,
		children: []*OpTree{},
	}
}

func (o *OpTree) Then(ot *OpTree) *OpTree {
	return o.ThenInParallel(ot)
}

func (o *OpTree) ThenDo(op Op) *OpTree {
	return o.Then(Do(op))
}

func (o *OpTree) ThenInParallel(ots ...*OpTree) *OpTree {
	if len(o.children) > 0 {
		panic("can only call ThenDo or ThenDoInParallel once per node.")
	}
	o.children = ots
	return o
}

func (o *OpTree) ThenDoInParallel(ops ...Op) *OpTree {
	return o.ThenInParallel(transformOpsToOpTrees(ops...)...)
}

func Sequentially(opts ...*OpTree) *OpTree {
	for i := len(opts) - 2; i >= 0; i-- {
		opts[i].Then(opts[i+1])
	}
	return opts[0]
}

func DoSequentially(ops ...Op) *OpTree {
	return Sequentially(transformOpsToOpTrees(ops...)...)
}

func InParallel(opts ...*OpTree) *OpTree {
	return DoNothing().ThenInParallel(opts...)
}

func DoInParallel(ops ...Op) *OpTree {
	return InParallel(transformOpsToOpTrees(ops...)...)
}

func transformOpsToOpTrees(ops ...Op) []*OpTree {
	ots := make([]*OpTree, len(ops))
	for i, op := range ops {
		ots[i] = Do(op)
	}
	return ots
}

func (o *OpTree) ToDebugString() string {
	return o.toDebugString(0)
}

func (o *OpTree) toDebugString(i int) string {
	r := strings.Repeat(" ", i)
	if i > 0 {
		r = r + "↳ "
	}
	if o.hasOp {
		r = r + o.op.ToDebugString()
	} else {
		r = r + "NoOp"
	}
	for _, c := range o.children {
		r += "\n" + c.toDebugString(i+1)
	}
	return r
}

func (o *OpTree) ValidateAll(errs *util.ErrList) {
	if o.hasOp {
		o.op.Validate(errs)
	}
	for _, c := range o.children {
		c.ValidateAll(errs)
	}
}

func (o *OpTree) Execute() []error {
	result := make(chan []error)
	go o.ExecuteAsync(result)
	return <-result
}

func (o *OpTree) ExecuteAsync(errorsToParent chan []error) {
	if o.hasOp {
		e := o.op.Execute()
		if e != nil {
			errorsToParent <- []error{e}
		}
	}
	errors := []error{}
	errorsFromChildren := make(chan []error)
	for _, c := range o.children {
		go c.ExecuteAsync(errorsFromChildren)
	}
	for _, _ = range o.children {
		errors = append(errors, (<-errorsFromChildren)...)
	}
	errorsToParent <- errors
}
