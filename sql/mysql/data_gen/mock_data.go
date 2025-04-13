package data_gen

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func Mock() error {

	// 链接数据库
	db := bootstrap.App().Mysql

	db.Select()

	_ = gofakeit.Seed(0)

	// 创建老师
	//	学生
	nowTime := time.Now().Unix()

}
