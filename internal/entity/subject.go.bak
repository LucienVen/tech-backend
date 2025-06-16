package entity

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	sq "github.com/Masterminds/squirrel"
)

func GetAllSubject() ([]Subject, error) {
	db := bootstrap.App.GetDB()

	subjects := make([]Subject, 0)
	sql, _, _ := sq.Select("*").From(SubjectTable).Where(sq.Eq{"is_delete": 0}).ToSql()
	//log.Info("sql:", zap.String("sql", sql))
	err := db.Select(&subjects, sql, 0)
	if err != nil {
		return subjects, err
	}
	return subjects, nil

}

// 获取指定学科信息
func GetSubjectByName(names ...string) ([]Subject, error) {
	db := bootstrap.App.GetDB()

	subjects := make([]Subject, 0)
	sql, args, _ := sq.Select("*").From(SubjectTable).Where(sq.Eq{"is_delete": 0}).Where(sq.Eq{"name": names}).ToSql()

	//log.Info("sql:", zap.String("sql", sql), zap.Any("names", names), zap.Any("args", args))

	err := db.Select(&subjects, sql, args...)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}
