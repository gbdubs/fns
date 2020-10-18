package util

import (
	"fmt"
)

type ErrList struct {
	errs []string
}

func NewErrList() *ErrList {
	return &ErrList{
		errs: []string{},
	}
}

func (e *ErrList) Appendf(f string, args ...interface{}) {
	e.errs = append(e.errs, fmt.Sprintf(f, args...))
}

func (e *ErrList) IsEmpty() bool {
	return len(e.errs) == 0
}

func (e *ErrList) String() string {
	s := ""
	for _, err := range e.errs {
		s = s + " * " + err + "\n"
	}
	return s
}
