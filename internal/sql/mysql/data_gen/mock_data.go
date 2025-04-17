package data_gen

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"time"
)

const (
	CreatorMock string = "mock"
	UpdaterMock string = "mock"
)

func Mock() error {

	// 链接数据库
	//db := bootstrap.App().Mysql

	_ = gofakeit.Seed(0)

	// 创建老师
	//	学生
	//nowTime := time.Now().Unix()

	insertSubjects()
	return nil
}

// insertTeachers 插入老师数据
func InsertTeachers() {
	//db := bootstrap.App().Mysql
	//teachers := []string{"张伟","王伟","王芳","李伟", "王秀英", "李秀英","李娜","张秀英","刘伟","张敏"}
	//
	//timenow := time.Now().Unix()
	//for _, name := range teachers {
	//	db.Exec("INSERT INTO teacher")
	//}
}

// insertSubjects 插入学科数据
func insertSubjects() {
	db := bootstrap.App().Mysql
	subjects := []string{"语文", "数学", "英语", "科学", "体育"}
	timenow := time.Now().Unix()
	for _, name := range subjects {

		res, err := db.Exec("INSERT INTO subject (name, description, is_delete, creator, updater, create_time, update_time) VALUES (?, ?, 0, ?, ?, ?, ?)",
			name, gofakeit.Sentence(5), CreatorMock, UpdaterMock, timenow, timenow)
		if err != nil {
			log.Info("insert subject failed.", zap.String("name", name), zap.Error(err))
		} else {
			log.Info("insert subject.", zap.String("name", name), zap.Any("res", res))
		}

	}
}
