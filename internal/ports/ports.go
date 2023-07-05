package ports

import (
	"context"
	"log"

	"github.com/xHappyface/school/internal/adapters/mysqlDB"
	"github.com/xHappyface/school/internal/core/courses"
	"github.com/xHappyface/school/internal/core/professors"
	"github.com/xHappyface/school/internal/core/students"
)

type CourseRepository interface {
	Create(*courses.Course) error
	ReadByID(id string) (*courses.Course, error)
	ReadByName(name string) (*courses.Course, error)
	Update(*courses.Course) error
	DeleteByID(id string) error
}

type ProfessorRepository interface {
	Create(*professors.Professor) error
	ReadByID(id string) (*professors.Professor, error)
	ReadByName(name string) (*professors.Professor, error)
	Update(*professors.Professor) error
	DeleteByID(id string) error
}

type StudentRepository interface {
	Create(*students.Student) error
	ReadByID(id string) (*students.Student, error)
	ReadByName(name string) (*students.Student, error)
	Update(*students.Student) error
	DeleteByID(id string) error
}

type SchoolService struct {
	DB            *mysqlDB.School
	CourseRepo    CourseRepository
	ProfessorRepo ProfessorRepository
	StudentRepo   StudentRepository
}

func NewSchoolService(ctx context.Context, l *log.Logger) (*SchoolService, error) {
	db, err := mysqlDB.NewSchoolDB(l)
	if err != nil {
		return new(SchoolService), err
	}
	return &SchoolService{
		DB:            db,
		CourseRepo:    mysqlDB.NewSQLCourseRepository(db, ctx, l),
		ProfessorRepo: mysqlDB.NewSQLProfessorRepository(db, ctx, l),
		StudentRepo:   mysqlDB.NewSQLStudentRepository(db, ctx, l),
	}, nil
}
