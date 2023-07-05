package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/xHappyface/school/internal/ports"
	"github.com/xHappyface/school/pkg/cli"

	"github.com/joho/godotenv"
)

func main() {
	l := log.New(os.Stderr, "SCHOOL: ", log.LstdFlags|log.Lmsgprefix|log.Lshortfile)
	// load environment
	if err := godotenv.Load(); err != nil {
		l.Fatalln("ERR:", err)
	}
	// set context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// logic:
	school, err := ports.NewSchoolService(ctx, l)
	if err != nil {
		l.Fatalln("ERR:", err)
	}
	defer school.DB.Close()
	cl := cli.NewCLIRepository(os.Stdin, os.Stdout, l)
	if err = cl.Run(school); err != nil {
		l.Fatalln("ERR:", err)
	}
}
