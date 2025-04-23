package utils

import (
	"encoding/json"
	"fmt"
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
)

func FirstInit() {
	app := bootstrap.Run()
	app.StartHeartbeat()
	defer app.CloseApplication()

	log.InitLogger(app.Env)
}

func StructPrintf(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
