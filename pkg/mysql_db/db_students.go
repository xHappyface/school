package mysql_db

import (
	"context"
	"database/sql"

	"github.com/xHappyface/school/api/students"
	"github.com/xHappyface/school/logger"
)

type SQLStudentRepository struct {
	db     *School
	ctx    context.Context
	logger *logger.SchoolLogger
}

func NewSQLStudentRepository(db *School, ctx context.Context, l *logger.SchoolLogger) *SQLStudentRepository {
	return &SQLStudentRepository{
		db:     db,
		ctx:    ctx,
		logger: l,
	}
}

func (repo *SQLStudentRepository) Create(cfg *students.Student) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, `insert into students(id, name, age, address, phone, if_international, if_on_probation)
										values (?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "student created")
	return nil
}

func (repo *SQLStudentRepository) ReadByID(id string) (*students.Student, error) {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from students where id=?;")
	if err != nil {
		return new(students.Student), err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, id)
	if err != nil {
		return new(students.Student), err
	}
	defer rows.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, "scanning rows...")
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "student retrieved")
	return student, nil
}

func (repo *SQLStudentRepository) ReadByName(name string) (*students.Student, error) {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from students where name=?;")
	if err != nil {
		return new(students.Student), err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, name)
	if err != nil {
		return new(students.Student), err
	}
	defer rows.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, "scanning rows...")
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "student retrieved")
	return student, nil
}

func (repo *SQLStudentRepository) Update(cfg *students.Student) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "update student set id=?, name=?, age=?, address=?, phone=?, if_on_probation=? where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "student updated")
	return nil
}

func (repo *SQLStudentRepository) DeleteByID(id string) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "delete from students where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "student deleted")
	return nil
}
