package entity

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func GetAllStudent() ([]Student, error) {
	db := bootstrap.App.GetDB()
	students := make([]Student, 0)

	sql, args, err := sq.Select("*").From(StudentTable).Where(sq.Eq{"is_delete": 0}).ToSql()
	if err != nil {
		log.Error("failed to build SQL", zap.Error(err))
		return nil, err
	}

	log.Info("executing SQL", zap.String("sql", sql), zap.Any("args", args))

	err = db.Select(&students, sql, args...)
	if err != nil {
		log.Error("failed to execute query", zap.Error(err))
		return nil, err
	}

	return students, nil
}
