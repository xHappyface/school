package courses

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var (
	errInvalidName = errors.New("invalid name")
)

type Course struct {
	ID   string
	Name string
}

func NewCourse(r io.Reader, w io.Writer) error {
	cfg, err := getCourseConfig(r, w)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "New course created.", cfg)
	return nil
}

func getCourseConfig(r io.Reader, w io.Writer) (*Course, error) {
	scanner := bufio.NewScanner(r)
	fmt.Fprint(w, "Enter course name: ")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return new(Course), err
	}
	name := strings.ToUpper(scanner.Text())
	match, err := regexp.MatchString(`^\b[ A-Z0-9]+\b$`, name)
	if err != nil {
		return new(Course), err
	}
	if !match {
		return new(Course), errInvalidName
	}
	course := &Course{
		ID:   uuid.NewString(),
		Name: name,
	}
	return course, nil
}
