package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

const (
	MysqlInterval = 30
)

var App Application

type Application struct {
	Env   *Env
	Mysql *sqlx.DB
}

func Run() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mysql = NewMysqlDatabase(app.Env)
	App = *app
	return App
}

func (app *Application) GetDB() *sqlx.DB {
	return app.Mysql
}

// 程序退出，清理资源
func (app *Application) CloseApplication() {
	if app.Mysql != nil {
		err := app.Mysql.Close()
		if err != nil {
			log.Fatal("close connect DB failed, err:", err)
		}

		log.Println("Connection to MysqlDB closed.")
	}
}

func (app *Application) StartHeartbeat() {
	app.startMysqlHeartbeat()
}

// 心跳检测
func (app *Application) startMysqlHeartbeat() {
	ticker := time.NewTicker(time.Second * MysqlInterval)
	go func() {
		for range ticker.C {
			if err := app.Mysql.Ping(); err != nil {
				log.Printf("数据库心跳失败：%v", err)
				// 可以考虑告警、重连等
			} else {
				log.Println("数据库心跳正常")
			}
		}
	}()
}
