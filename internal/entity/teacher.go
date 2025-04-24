package entity

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func GetAllTeacher() ([]Teacher, error) {
	db := bootstrap.App.Mysql
	teachers := make([]Teacher, 0)

	// 明确指定需要的列
	sql, args, err := sq.Select("*").From(TeacherTable).Where(sq.Eq{"is_delete": 0}).ToSql()
	if err != nil {
		log.Error("failed to build SQL", zap.Error(err))
		return nil, err
	}

	log.Info("executing SQL", zap.String("sql", sql), zap.Any("args", args))
	err = db.Select(&teachers, sql, args...)
	if err != nil {
		log.Error("failed to execute query", zap.Error(err))
		return nil, err
	}

	return teachers, nil
}

// -----------------------------------
type TeacherEntity struct {
	db      *sqlx.DB
	teacher Teacher
}

func NewTeacherEntity(db *sqlx.DB, teacher Teacher) *TeacherEntity {
	return &TeacherEntity{
		teacher: teacher,
		db:      db,
	}

}

func (t *TeacherEntity) GetId() string {
	return t.teacher.Id
}

// 更新责任学科信息
func (t *TeacherEntity) UpdateSubjectId(subjectId, handler string) error {

	updateData := t.teacher.GenWhenUpdateMap(handler)

	sql, args, err := sq.Update(TeacherTable).
		Set("subject_id", subjectId).
		SetMap(updateData).Where(sq.Eq{"id": t.GetId()}).ToSql()

	if err != nil {
		return err
	}

	log.Info("UpdateSubjectId executing SQL", zap.String("sql", sql), zap.Any("args", args))

	_, err = t.db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
