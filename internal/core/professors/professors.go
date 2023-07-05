package professors

import (
	"fmt"
	"io"
)

type Professor struct {
	ID              string
	Name            string
	Age             uint8
	Address         string
	Phone           uint
	Salary          float64
	IfReceivedBonus bool
}

func NewProfessor(r io.Reader, w io.Writer) error {
	fmt.Fprintln(w, "New professor created.")
	return nil
}
