package cli

import (
	"_school/internal/ports"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

var (
	wrnExtraStatementsTruncated = errors.New("extra statement(s) truncated")

	errExitSignal = errors.New("exit signal")
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

func (repo *CLIRepository) Run(sch *ports.SchoolService) error {
	fmt.Fprintln(repo.Writer, "Welcome.")
	scanner := bufio.NewScanner(repo.Reader)
	var input bytes.Buffer
	for {
		fmt.Fprint(repo.Writer, "> ")
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
			log.Println("WRN:", wrnExtraStatementsTruncated)
		}
		args := strings.Fields(input.String())
		input.Reset()
		if err := repo.execute(sch, args); err != nil && !errors.Is(err, errExitSignal) {
			return err
		} else if errors.Is(err, errExitSignal) {
			fmt.Fprintln(repo.Writer, "Goodbye!")
			break
		}
	}
	return nil
}

func (repo *CLIRepository) execute(sch *ports.SchoolService, args []string) error {
	fmt.Fprintln(repo.Writer, args)
	return nil
}
