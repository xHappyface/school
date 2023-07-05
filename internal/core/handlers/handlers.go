package handlers

import (
	"errors"
	"io"

	"github.com/xHappyface/school/internal/core/courses"
	"github.com/xHappyface/school/internal/core/professors"
	"github.com/xHappyface/school/internal/core/students"
	"github.com/xHappyface/school/internal/ports"
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
