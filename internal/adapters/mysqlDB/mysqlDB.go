package mysqlDB

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/xHappyface/school/internal/core/courses"
	"github.com/xHappyface/school/internal/core/professors"
	"github.com/xHappyface/school/internal/core/students"

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

func NewSchoolDB(l *log.Logger) (*School, error) {
	db, err := connect(l)
	if err != nil {
		return new(School), err
	}
	return &School{school: db}, nil
}

type SQLCourseRepository struct {
	db     *School
	ctx    context.Context
	logger *log.Logger
}

func NewSQLCourseRepository(db *School, ctx context.Context, l *log.Logger) *SQLCourseRepository {
	return &SQLCourseRepository{
		db:     db,
		ctx:    ctx,
		logger: l,
	}
}

type SQLProfessorRepository struct {
	db     *School
	ctx    context.Context
	logger *log.Logger
}

func NewSQLProfessorRepository(db *School, ctx context.Context, l *log.Logger) *SQLProfessorRepository {
	return &SQLProfessorRepository{
		db:     db,
		ctx:    ctx,
		logger: l,
	}
}

type SQLStudentRepository struct {
	db     *School
	ctx    context.Context
	logger *log.Logger
}

func NewSQLStudentRepository(db *School, ctx context.Context, l *log.Logger) *SQLStudentRepository {
	return &SQLStudentRepository{
		db:     db,
		ctx:    ctx,
		logger: l,
	}
}

func connect(l *log.Logger) (*sql.DB, error) {
	l.Println("connecting to sql database...")
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
	l.Println("connected to database")
	return db, nil
}

func (schoolDB *School) Close() error {
	return schoolDB.school.Close()
}

func (repo *SQLCourseRepository) Create(cfg *courses.Course) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "insert into courses(id, name) values (?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, cfg.ID, cfg.Name)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("course created")
	return nil
}

func (repo *SQLCourseRepository) ReadByID(id string) (*courses.Course, error) {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from courses where id=?;")
	if err != nil {
		return new(courses.Course), err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, id)
	if err != nil {
		return new(courses.Course), err
	}
	defer rows.Close()
	repo.logger.Println("scanning rows...")
	course := new(courses.Course)
	for rows.Next() {
		err = rows.Scan(&course.ID, &course.Name)
		if err != nil {
			return new(courses.Course), err
		}
	}
	if err = rows.Err(); err != nil {
		return new(courses.Course), err
	}
	if course.ID == "" {
		return new(courses.Course), errZeroRowsRetrieved
	}
	repo.logger.Println("course retrieved")
	return course, nil
}

func (repo *SQLCourseRepository) ReadByName(name string) (*courses.Course, error) {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from courses where name=?;")
	if err != nil {
		return new(courses.Course), err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, name)
	if err != nil {
		return new(courses.Course), err
	}
	defer rows.Close()
	repo.logger.Println("scanning rows...")
	course := new(courses.Course)
	for rows.Next() {
		err = rows.Scan(&course.ID, &course.Name)
		if err != nil {
			return new(courses.Course), err
		}
	}
	if err = rows.Err(); err != nil {
		return new(courses.Course), err
	}
	if course.ID == "" {
		return new(courses.Course), errZeroRowsRetrieved
	}
	repo.logger.Println("course retrieved")
	return course, nil
}

func (repo *SQLCourseRepository) Update(cfg *courses.Course) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "update courses set id=?, name=? where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, cfg.ID, cfg.Name, cfg.ID)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("course updated")
	return nil
}

func (repo *SQLCourseRepository) DeleteByID(id string) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "delete from courses where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, id)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("course deleted")
	return nil
}

func (repo *SQLProfessorRepository) Create(cfg *professors.Professor) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, `insert into professors(id, name, age, address, phone, salary, if_received_bonus)
										values (?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, cfg.ID, cfg.Name, cfg.Age, cfg.Address, cfg.Phone, cfg.Salary, cfg.IfReceivedBonus)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("professor created")
	return nil
}

func (repo *SQLProfessorRepository) ReadByID(id string) (*professors.Professor, error) {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from professors where id=?;")
	if err != nil {
		return new(professors.Professor), err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, id)
	if err != nil {
		return new(professors.Professor), err
	}
	defer rows.Close()
	repo.logger.Println("scanning rows...")
	professor := new(professors.Professor)
	for rows.Next() {
		err = rows.Scan(&professor.ID, &professor.Name, &professor.Age, &professor.Address, &professor.Phone, &professor.Salary, &professor.IfReceivedBonus)
		if err != nil {
			return new(professors.Professor), err
		}
	}
	if err = rows.Err(); err != nil {
		return new(professors.Professor), err
	}
	if professor.ID == "" {
		return new(professors.Professor), errZeroRowsRetrieved
	}
	repo.logger.Println("professor retrieved")
	return professor, nil
}

func (repo *SQLProfessorRepository) ReadByName(name string) (*professors.Professor, error) {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from professors where name=?;")
	if err != nil {
		return new(professors.Professor), err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, name)
	if err != nil {
		return new(professors.Professor), err
	}
	defer rows.Close()
	repo.logger.Println("scanning rows...")
	professor := new(professors.Professor)
	for rows.Next() {
		err = rows.Scan(&professor.ID, &professor.Name, &professor.Age, &professor.Address, &professor.Phone, &professor.Salary, &professor.IfReceivedBonus)
		if err != nil {
			return new(professors.Professor), err
		}
	}
	if err = rows.Err(); err != nil {
		return new(professors.Professor), err
	}
	if professor.ID == "" {
		return new(professors.Professor), errZeroRowsRetrieved
	}
	repo.logger.Println("professor retrieved")
	return professor, nil
}

func (repo *SQLProfessorRepository) Update(cfg *professors.Professor) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "update professors set id=?, name=?, age=?, address=?, phone=?, salary=? where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, cfg.ID, cfg.Name, cfg.Age, cfg.Address, cfg.Phone, cfg.Salary, cfg.ID)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("professor updated")
	return nil
}

func (repo *SQLProfessorRepository) DeleteByID(id string) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "delete from professors where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, id)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("professor deleted")
	return nil
}

func (repo *SQLStudentRepository) Create(cfg *students.Student) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, `insert into students(id, name, age, address, phone, if_international, if_on_probation)
										values (?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, cfg.ID)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("student created")
	return nil
}

func (repo *SQLStudentRepository) ReadByID(id string) (*students.Student, error) {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from students where id=?;")
	if err != nil {
		return new(students.Student), err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, id)
	if err != nil {
		return new(students.Student), err
	}
	defer rows.Close()
	repo.logger.Println("scanning rows...")
	student := new(students.Student)
	for rows.Next() {
		err = rows.Scan(&student.ID, &student.Name, &student.Age, &student.Address, &student.Phone, &student.IfInternational, &student.IfOnProbation)
		if err != nil {
			return new(students.Student), err
		}
	}
	if err = rows.Err(); err != nil {
		return new(students.Student), err
	}
	if student.ID == "" {
		return new(students.Student), errZeroRowsRetrieved
	}
	repo.logger.Println("student retrieved")
	return student, nil
}

func (repo *SQLStudentRepository) ReadByName(name string) (*students.Student, error) {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from students where name=?;")
	if err != nil {
		return new(students.Student), err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, name)
	if err != nil {
		return new(students.Student), err
	}
	defer rows.Close()
	repo.logger.Println("scanning rows...")
	student := new(students.Student)
	for rows.Next() {
		err = rows.Scan(&student.ID, &student.Name, &student.Age, &student.Address, &student.Phone, &student.IfInternational, &student.IfOnProbation)
		if err != nil {
			return new(students.Student), err
		}
	}
	if err = rows.Err(); err != nil {
		return new(students.Student), err
	}
	if student.ID == "" {
		return new(students.Student), errZeroRowsRetrieved
	}
	repo.logger.Println("student retrieved")
	return student, nil
}

func (repo *SQLStudentRepository) Update(cfg *students.Student) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "update student set id=?, name=?, age=?, address=?, phone=?, if_on_probation=? where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, cfg.ID, cfg.Name, cfg.Age, cfg.Address, cfg.Phone, cfg.IfOnProbation, cfg.ID)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("student updated")
	return nil
}

func (repo *SQLStudentRepository) DeleteByID(id string) error {
	repo.logger.Println(LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "delete from students where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Println(LOG_EXECUTING_STMT)
	var result sql.Result
	result, err = stmt.ExecContext(repo.ctx, id)
	if err != nil {
		return err
	}
	var affected int64
	affected, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if !(affected > 0) {
		return errZeroRowsAffected
	}
	repo.logger.Println("student deleted")
	return nil
}
