package entity

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// 获取所有班级信息
func GetAllClass() ([]Class, error) {
	app := bootstrap.Run()
	db := app.GetDB()

	classes := make([]Class, 0)

	// 明确指定需要的列
	sql, args, err := sq.Select("*").From(ClassTable).Where(sq.Eq{"is_delete": 0}).ToSql()
	if err != nil {
		log.Error("failed to build SQL", zap.Error(err))
		return nil, err
	}

	log.Info("executing SQL", zap.String("sql", sql), zap.Any("args", args))
	err = db.Select(&classes, sql, args...)
	if err != nil {
		log.Error("failed to execute query", zap.Error(err))
		return nil, err
	}

	return classes, nil
}

// ------------------------
type ClassEntity struct {
	db    *sqlx.DB
	class Class
}

func NewClassEntity(db *sqlx.DB, class Class) *ClassEntity {
	return &ClassEntity{
		class: class,
		db:    db,
	}
}

func (ce *ClassEntity) GetId() string {
	return ce.class.Id
}

func (ce *ClassEntity) UpdateMainTeacherId(teacherId string, handler string) error {
	updateData := ce.class.GenWhenUpdateMap(handler)

	sql, args, err := sq.Update(ClassTable).
		Set("main_teacher_id", teacherId).
		SetMap(updateData).Where(sq.Eq{"id": ce.GetId()}).ToSql()

	if err != nil {
		return err
	}

	log.Info("UpdateMainTeacherId sql:", zap.String("sql", sql), zap.Any("args", args))

	_, err = ce.db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
