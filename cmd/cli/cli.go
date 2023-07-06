package cli

import (
	"errors"
	"io"

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
