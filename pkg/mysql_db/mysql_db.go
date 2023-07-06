package mysql_db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/xHappyface/school/logger"

	_ "github.com/go-sql-driver/mysql"
)

const (
	LOG_PREPARING_STMT = "preparing sql statement..."
	LOG_EXECUTING_STMT = "executing sql statement..."
)

var (
	errZeroRowsAffected  = errors.New("zero rows affected")
	errZeroRowsRetrieved = errors.New("zero rows retrieved")
)

type School struct {
	school *sql.DB
}

func NewSchoolDB(l *logger.SchoolLogger) (*School, error) {
	db, err := connect(l)
	if err != nil {
		return new(School), err
	}
	return &School{school: db}, nil
}

func connect(l *logger.SchoolLogger) (*sql.DB, error) {
	l.Log(logger.LOG_LEVEL_INFO, "connecting to sql database...")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	source := fmt.Sprintf("%s:%s@tcp(localhost:3306)/school", user, pass)
	db, err := sql.Open("mysql", source)
	if err != nil {
		return new(sql.DB), err
	}
	if err = db.Ping(); err != nil {
		return new(sql.DB), err
	}
	l.Log(logger.LOG_LEVEL_INFO, "connected to database")
	return db, nil
}

func (schoolDB *School) Close() error {
	return schoolDB.school.Close()
}
