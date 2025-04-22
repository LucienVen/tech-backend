package entity

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	sq "github.com/Masterminds/squirrel"
	"go.uber.org/zap"
)

func GetAllSubject() ([]Subject, error) {
	db := bootstrap.App().Mysql
	subjects := make([]Subject, 0)
	sql, _, _ := sq.Select("*").From(SubjectTable).Where(sq.Eq{"is_delete": 0}).ToSql()
	log.Info("sql:", zap.String("sql", sql))
	err := db.Select(&subjects, sql, 0)
	if err != nil {
		return subjects, err
	}
	return subjects, nil

}
