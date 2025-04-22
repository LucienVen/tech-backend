package utils

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
)

func FirstInit() {
	app := bootstrap.App()
	app.StartHeartbeat()
	defer app.CloseApplication()

	log.InitLogger(app.Env)
}
