package utils

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/pkg/log"
)

func FirstInit() func() {
	app := bootstrap.Run()
	app.StartHeartbeat()

	//log2.Printf("%+v", app.Mysql)
	//log2.Printf("%+v", bootstrap.App.Mysql)
	//log2.Printf("%+v", bootstrap.App.GetDB())

	log.InitLogger()

	return app.CloseApplication
}

func StructPrintf(v interface{}) {
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(b))
}

// 浮点数保留 N 位小数(返回 float) 不四舍五入
func FormatFloat2Float(num float64, decimal int) float64 {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		d = math.Pow10(decimal)
	}

	return math.Trunc(num*d) / d
}
