package main

import (
	"fmt"
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
)

func main() {
	fmt.Println("hello, world")

	app := bootstrap.Run()
	app.StartHeartbeat()
	defer app.CloseApplication()

	// 初始化日志组件
	log.InitLogger(app.Env)
	defer log.Sync()

	select {}
}
