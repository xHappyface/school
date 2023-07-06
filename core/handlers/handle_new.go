package handlers

import (
	"github.com/xHappyface/school/api/professors"
	"github.com/xHappyface/school/api/students"
	"github.com/xHappyface/school/pkg/cli"
)

func (handler *SchoolHandler) HandleCmdNew() error {
	var err error
	switch handler.obj {
	case "course":
		if err = cli.NewCourse(handler.r, handler.w, handler.sch.CourseRepo); err != nil {
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
