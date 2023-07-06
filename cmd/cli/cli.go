package cli

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/xHappyface/school/api/ports"
	"github.com/xHappyface/school/core/handlers"
	"github.com/xHappyface/school/logger"
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
	Logger *logger.SchoolLogger
}

func NewCLIRepository(r io.Reader, w io.Writer, l *logger.SchoolLogger) *CLIRepository {
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
			cl.Logger.Log(logger.LOG_LEVEL_WRN, wrnExtraStatementsTruncated.Error())
		}
		args := strings.Fields(input.String())
		input.Reset()
		if !(len(args) > 0) {
			cl.Logger.Log(logger.LOG_LEVEL_ERR, errTooFewArgs.Error())
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
		cl.Logger.Log(logger.LOG_LEVEL_ERR, errTooFewArgs.Error())
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
			cl.Logger.Log(logger.LOG_LEVEL_ERR, err.Error())
		}
	default:
		cl.Logger.Log(logger.LOG_LEVEL_ERR, errInvalidCommand.Error())
	}
	return nil
}
