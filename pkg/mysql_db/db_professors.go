package mysql_db

import (
	"context"
	"database/sql"

	"github.com/xHappyface/school/api/professors"
	"github.com/xHappyface/school/logger"
)

type SQLProfessorRepository struct {
	db     *School
	ctx    context.Context
	logger *logger.SchoolLogger
}

func NewSQLProfessorRepository(db *School, ctx context.Context, l *logger.SchoolLogger) *SQLProfessorRepository {
	return &SQLProfessorRepository{
		db:     db,
		ctx:    ctx,
		logger: l,
	}
}

func (repo *SQLProfessorRepository) Create(cfg *professors.Professor) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, `insert into professors(id, name, age, address, phone, salary, if_received_bonus)
										values (?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "professor created")
	return nil
}

func (repo *SQLProfessorRepository) ReadByID(id string) (*professors.Professor, error) {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from professors where id=?;")
	if err != nil {
		return new(professors.Professor), err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, id)
	if err != nil {
		return new(professors.Professor), err
	}
	defer rows.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, "scanning rows...")
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "professor retrieved")
	return professor, nil
}

func (repo *SQLProfessorRepository) ReadByName(name string) (*professors.Professor, error) {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from professors where name=?;")
	if err != nil {
		return new(professors.Professor), err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, name)
	if err != nil {
		return new(professors.Professor), err
	}
	defer rows.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, "scanning rows...")
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "professor retrieved")
	return professor, nil
}

func (repo *SQLProfessorRepository) Update(cfg *professors.Professor) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "update professors set id=?, name=?, age=?, address=?, phone=?, salary=? where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "professor updated")
	return nil
}

func (repo *SQLProfessorRepository) DeleteByID(id string) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "delete from professors where id=?;")
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "professor deleted")
	return nil
}
