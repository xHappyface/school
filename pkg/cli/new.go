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
)

var (
	errInvalidName = errors.New("invalid name")
)

func NewCourse(r io.Reader, w io.Writer) error {
	cfg, err := getCourseConfig(r, w)
	if err != nil {
		return err
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
		return new(courses.Course), errInvalidName
	}
	course := &courses.Course{
		ID:   uuid.NewString(),
		Name: name,
	}
	return course, nil
}
