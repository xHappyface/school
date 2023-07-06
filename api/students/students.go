package students

import (
	"fmt"
	"io"
)

type Student struct {
	ID              string
	Name            string
	Age             uint8
	Address         string
	Phone           uint
	IfInternational bool
	IfOnProbation   bool
}

func NewStudent(r io.Reader, w io.Writer) error {
	fmt.Fprintln(w, "New student created.")
	return nil
}
