package main

import (
	"context"
	"os"
	"time"

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
	// set context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// logic:
	school, err := ports.NewSchoolService(ctx, l)
	if err != nil {
		l.Log(logger.LOG_LEVEL_FATAL_ERR, err.Error())
	}
	defer school.DB.Close()
	cl := cli.NewCLIRepository(os.Stdin, os.Stdout, l)
	if err = cl.Run(school); err != nil {
		l.Log(logger.LOG_LEVEL_FATAL_ERR, err.Error())
	}
}
