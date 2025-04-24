package entity

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func GetAllExams() ([]Exam, error) {
	db := bootstrap.App.GetDB()

	exams := make([]Exam, 0)
	sql, args, _ := sq.Select("*").From(ExamTable).Where(sq.Eq{"is_delete": 0}).ToSql()
	log.Info("GetAllExams sql:", zap.String("sql", sql), zap.Any("args", args))

	err := db.Select(&exams, sql, args...)
	if err != nil {
		log.Error("failed to execute query", zap.Error(err))
		return nil, err
	}

	return exams, nil
}
