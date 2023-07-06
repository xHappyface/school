package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/xHappyface/school/api/courses"
	"github.com/xHappyface/school/api/ports"
	"github.com/xHappyface/school/pkg/mysql_db"
)

func NewCourse(r io.Reader, w io.Writer, repo ports.CourseRepository) error {
	cfg, err := getCourseConfig(r, w)
	if err != nil {
		return err
	}
	if _, err = repo.ReadByID(cfg.ID); !(errors.Is(err, mysql_db.ErrZeroRowsRetrieved)) {
		fmt.Fprintln(w, err)
		return ErrObjectAlreadyExists
	}
	if _, err = repo.ReadByName(cfg.Name); !(errors.Is(err, mysql_db.ErrZeroRowsRetrieved)) {
		fmt.Fprintln(w, err)
		return ErrObjectAlreadyExists
	}
	fmt.Fprintln(w, "New course created.", cfg)
	return nil
}

func getCourseConfig(r io.Reader, w io.Writer) (*courses.Course, error) {
	scanner := bufio.NewScanner(r)
	fmt.Fprint(w, "Enter course name: ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return new(courses.Course), err
	}
	name := strings.ToUpper(scanner.Text())
	match, err := regexp.MatchString(`^\b[ A-Z0-9]+\b$`, name)
	if err != nil {
		return new(courses.Course), err
	}
	if !match {
		return new(courses.Course), ErrInvalidName
	}
	course := &courses.Course{
		ID:   uuid.NewString(),
		Name: name,
	}
	return course, nil
}
