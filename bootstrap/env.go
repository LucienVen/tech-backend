package bootstrap

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Env struct {
	AppEnv  string `mapstructure:"APP_ENV"`
	AppName string `mapstructure:"APP_NAME"`
	DBHost  string `mapstructure:"DB_HOST"`
	DBPort  string `mapstructure:"DB_PORT"`
	DBUser  string `mapstructure:"DB_USER"`
	DBPass  string `mapstructure:"DB_PASS"`
	DBName  string `mapstructure:"DB_NAME"`
}

// GetGinMode 根据环境变量返回对应的 Gin 模式
func (e *Env) GetGinMode() string {
	switch e.AppEnv {
	case "production":
		return gin.ReleaseMode
	case "testing":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv == "development" {
		log.Println("The Run is running in development env")
	}

	return &env
}
