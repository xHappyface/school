package ports

import (
	"context"

	"github.com/xHappyface/school/api/courses"
	"github.com/xHappyface/school/api/professors"
	"github.com/xHappyface/school/api/students"
	"github.com/xHappyface/school/logger"
	"github.com/xHappyface/school/pkg/mysql_db"
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
	DB            *mysql_db.School
	CourseRepo    CourseRepository
	ProfessorRepo ProfessorRepository
	StudentRepo   StudentRepository
}

func NewSchoolService(ctx context.Context, l *logger.SchoolLogger) (*SchoolService, error) {
	db, err := mysql_db.NewSchoolDB(l)
	if err != nil {
		return new(SchoolService), err
	}
	return &SchoolService{
		DB:            db,
		CourseRepo:    mysql_db.NewSQLCourseRepository(db, ctx, l),
		ProfessorRepo: mysql_db.NewSQLProfessorRepository(db, ctx, l),
		StudentRepo:   mysql_db.NewSQLStudentRepository(db, ctx, l),
	}, nil
}
