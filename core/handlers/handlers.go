package handlers

import (
	"errors"
	"io"

	"github.com/xHappyface/school/api/ports"
	"github.com/xHappyface/school/core/app/courses"
	"github.com/xHappyface/school/core/app/professors"
	"github.com/xHappyface/school/core/app/students"
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
		if err = courses.NewCourse(handler.r, handler.w); err != nil {
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
