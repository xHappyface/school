package mysql_db

import (
	"context"
	"database/sql"

	"github.com/xHappyface/school/api/courses"
	"github.com/xHappyface/school/logger"
)

type SQLCourseRepository struct {
	db     *School
	ctx    context.Context
	logger *logger.SchoolLogger
}

func NewSQLCourseRepository(db *School, ctx context.Context, l *logger.SchoolLogger) *SQLCourseRepository {
	return &SQLCourseRepository{
		db:     db,
		ctx:    ctx,
		logger: l,
	}
}

func (repo *SQLCourseRepository) Create(cfg *courses.Course) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "insert into courses(id, name) values (?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "course created")
	return nil
}

func (repo *SQLCourseRepository) ReadByID(id string) (*courses.Course, error) {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from courses where id=?;")
	if err != nil {
		return new(courses.Course), err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, id)
	if err != nil {
		return new(courses.Course), err
	}
	defer rows.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, "scanning rows...")
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "course retrieved")
	return course, nil
}

func (repo *SQLCourseRepository) ReadByName(name string) (*courses.Course, error) {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "select * from courses where name=?;")
	if err != nil {
		return new(courses.Course), err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
	var rows *sql.Rows
	rows, err = stmt.QueryContext(repo.ctx, name)
	if err != nil {
		return new(courses.Course), err
	}
	defer rows.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, "scanning rows...")
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "course retrieved")
	return course, nil
}

func (repo *SQLCourseRepository) Update(cfg *courses.Course) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "update courses set id=?, name=? where id=?;")
	if err != nil {
		return err
	}
	defer stmt.Close()
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_EXECUTING_STMT)
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "course updated")
	return nil
}

func (repo *SQLCourseRepository) DeleteByID(id string) error {
	repo.logger.Log(logger.LOG_LEVEL_INFO, LOG_PREPARING_STMT)
	stmt, err := repo.db.school.PrepareContext(repo.ctx, "delete from courses where id=?;")
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
	repo.logger.Log(logger.LOG_LEVEL_INFO, "course deleted")
	return nil
}
