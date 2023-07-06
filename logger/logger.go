package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	PREFIX_APP  string = "SCHOOL:"
	PREFIX_INFO string = "INFO"
	PREFIX_WRN  string = "WRN"
	PREFIX_ERR  string = "ERR"

	LOG_LEVEL_INFO uint8 = iota + 1
	LOG_LEVEL_WRN
	LOG_LEVEL_ERR
	LOG_LEVEL_NOTABLE_WRN
	LOG_LEVEL_FATAL_ERR
)

var (
	SchoolLog *SchoolLogger

	errInvalidLogLevel = errors.New("invalid log level")
)

type SchoolLogger struct {
	stdout *log.Logger
	stderr *log.Logger
}

func New() *SchoolLogger {
	return &SchoolLogger{
		stdout: log.New(os.Stdout, PREFIX_APP, log.LstdFlags|log.Lmsgprefix),
		stderr: log.New(os.Stderr, PREFIX_APP, log.LstdFlags|log.Lmsgprefix|log.Lshortfile),
	}
}

func (logger *SchoolLogger) Log(lv uint8, msg string) {
	switch lv {
	case LOG_LEVEL_INFO:
		logger.stdout.Println(fmt.Sprintf("%s: %s", PREFIX_INFO, msg))
	case LOG_LEVEL_WRN:
		logger.stdout.Println(fmt.Errorf("%s: %s", PREFIX_WRN, msg))
	case LOG_LEVEL_ERR:
		logger.stdout.Println(fmt.Errorf("%s: %s", PREFIX_ERR, msg))
	case LOG_LEVEL_NOTABLE_WRN:
		logger.stderr.Println(fmt.Errorf("%s: %s", PREFIX_WRN, msg))
	case LOG_LEVEL_FATAL_ERR:
		logger.stderr.Fatalln(fmt.Errorf("%s: %s", PREFIX_ERR, msg))
	}
}
