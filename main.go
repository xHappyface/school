package main

import (
	"context"
	"log"
	"os"
	"time"

	"_school/internal/ports"
	"_school/pkg/cli"

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
	c := cli.NewCLIRepository(os.Stdin, os.Stdout, l)
	if err = c.Run(school); err != nil {
		l.Fatalln("ERR:", err)
	}
}