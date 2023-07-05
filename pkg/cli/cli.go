package cli

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/xHappyface/school/internal/core/handlers"
	"github.com/xHappyface/school/internal/ports"
)

var (
	wrnExtraStatementsTruncated = errors.New("extra statement(s) truncated")

	errExitSignal     = errors.New("exit signal")
	errTooFewArgs     = errors.New("too few args")
	errInvalidCommand = errors.New("invalid command")
)

type CLIRepository struct {
	Reader io.Reader
	Writer io.Writer
	Logger *log.Logger
}

func NewCLIRepository(r io.Reader, w io.Writer, l *log.Logger) *CLIRepository {
	return &CLIRepository{
		Reader: r,
		Writer: w,
		Logger: l,
	}
}

func (cl *CLIRepository) Run(sch *ports.SchoolService) error {
	fmt.Fprintln(cl.Writer, "Welcome.")
	scanner := bufio.NewScanner(cl.Reader)
	var input bytes.Buffer
	for {
		fmt.Fprint(cl.Writer, "> ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return err
		}
		text := strings.ToLower(scanner.Text())
		if !(strings.ContainsRune(text, ';')) {
			input.WriteString(text + " ")
			continue
		}
		statements := strings.Split(text, ";")
		text = strings.ToLower(statements[0])
		input.WriteString(text)
		if len(strings.Fields(statements[1])) != 0 {
			cl.Logger.Println("WRN:", wrnExtraStatementsTruncated)
		}
		args := strings.Fields(input.String())
		input.Reset()
		if !(len(args) > 0) {
			cl.Logger.Println(errTooFewArgs)
			continue
		}
		if err := cl.execute(sch, args); err != nil && !errors.Is(err, errExitSignal) {
			return err
		} else if errors.Is(err, errExitSignal) {
			fmt.Fprintln(cl.Writer, "Goodbye!")
			break
		}
	}
	return nil
}

func (cl *CLIRepository) execute(sch *ports.SchoolService, args []string) error {
	if args[0] == "exit" {
		return errExitSignal
	}
	if !(len(args) >= 2) {
		cl.Logger.Println(errTooFewArgs)
		return nil
	}
	cmd := args[0]
	obj := args[1]
	if len(args) > 2 {
		args = args[2:]
	} else {
		args = []string{}
	}
	handler := handlers.NewSchoolHandler(cl.Reader, cl.Writer, sch, obj, args)
	var err error
	switch cmd {
	case "new":
		if err = handler.HandleCmdNew(); err != nil {
			cl.Logger.Println(err)
		}
	default:
		cl.Logger.Println(errInvalidCommand)
	}
	return nil
}
