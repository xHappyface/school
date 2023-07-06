package main

import (
	"os"

	"github.com/xHappyface/school/api/ports"
	"github.com/xHappyface/school/cmd/cli"
	"github.com/xHappyface/school/logger"

	"github.com/joho/godotenv"
)

func main() {
	l := logger.New()
	// load environment
	if err := godotenv.Load(); err != nil {
		l.Log(logger.LOG_LEVEL_FATAL_ERR, err.Error())
	}
	// logic:
	school, err := ports.NewSchoolService(l, 10_000)
	if err != nil {
		l.Log(logger.LOG_LEVEL_FATAL_ERR, err.Error())
	}
	defer school.DB.Close()
	cl := cli.NewCLIRepository(os.Stdin, os.Stdout, l)
	if err = cl.Run(school); err != nil {
		l.Log(logger.LOG_LEVEL_FATAL_ERR, err.Error())
	}
}
