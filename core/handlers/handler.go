package handlers

import (
	"errors"
	"io"

	"github.com/xHappyface/school/api/ports"
)

var (
	errInvalidObject = errors.New("invalid object")
)

type SchoolHandler struct {
	r    io.Reader
	w    io.Writer
	sch  *ports.SchoolService
	obj  string
	args []string
}

func NewSchoolHandler(r io.Reader, w io.Writer, sch *ports.SchoolService, obj string, args []string) *SchoolHandler {
	return &SchoolHandler{
		r:    r,
		w:    w,
		sch:  sch,
		obj:  obj,
		args: args,
	}
}
