package utils

import (
	"encoding/json"
	"fmt"
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/manager/log"
)

func FirstInit() func() {
	app := bootstrap.Run()
	app.StartHeartbeat()

	//log2.Printf("%+v", app.Mysql)
	//log2.Printf("%+v", bootstrap.App.Mysql)
	//log2.Printf("%+v", bootstrap.App.GetDB())

	log.InitLogger(app.Env)

	return app.CloseApplication
}

func StructPrintf(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}
