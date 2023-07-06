package handlers

import (
	"errors"
	"io"

	"github.com/xHappyface/school/api/ports"
	"github.com/xHappyface/school/api/professors"
	"github.com/xHappyface/school/api/students"
	"github.com/xHappyface/school/pkg/cli"
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

func (handler *SchoolHandler) HandleCmdNew() error {
	var err error
	switch handler.obj {
	case "course":
		if err = cli.NewCourse(handler.r, handler.w); err != nil {
			return err
		}
	case "professor":
		if err = professors.NewProfessor(handler.r, handler.w); err != nil {
			return err
		}
	case "student":
		if err = students.NewStudent(handler.r, handler.w); err != nil {
			return err
		}
	default:
		return errInvalidObject
	}
	return nil
}
