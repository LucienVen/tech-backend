package bootstrap

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	MysqlInterval = 30
)

// Application 应用程序核心结构
type Application struct {
	env   *Env
	mysql *sqlx.DB
	ctx   context.Context
}

// NewApplication 创建新的应用实例
func NewApplication() *Application {
	ctx := context.Background()
	return &Application{
		ctx: ctx,
	}
}

// Initialize 初始化应用
func (app *Application) Initialize() error {
	// 初始化环境配置
	app.env = NewEnv()

	// 初始化数据库连接
	app.mysql = NewMysqlDatabase(app.env)

	// 启动心跳检测
	app.startMysqlHeartbeat()

	return nil
}

// GetDB 获取数据库连接
func (app *Application) GetDB() *sqlx.DB {
	return app.mysql
}

// Close 清理资源
func (app *Application) Close() error {
	if app.mysql != nil {
		if err := app.mysql.Close(); err != nil {
			log.Printf("关闭数据库连接失败: %v", err)
			return err
		}
		log.Println("数据库连接已关闭")
	}
	return nil
}

// startMysqlHeartbeat 启动数据库心跳检测
func (app *Application) startMysqlHeartbeat() {
	ticker := time.NewTicker(time.Second * MysqlInterval)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := app.mysql.Ping(); err != nil {
					log.Printf("数据库心跳检测失败: %v", err)
					// TODO: 实现重连逻辑
					if err := app.reconnectDB(); err != nil {
						log.Printf("数据库重连失败: %v", err)
					}
				} else {
					log.Println("数据库心跳检测正常")
				}
			case <-app.ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

// reconnectDB 重新连接数据库
func (app *Application) reconnectDB() error {
	if app.mysql != nil {
		app.mysql.Close()
	}

	app.mysql = NewMysqlDatabase(app.env)
	return nil
}
