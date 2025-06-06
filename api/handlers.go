package api

import (
	"net/http"

	"github.com/LucienVen/tech-backend/internal/db"
	"github.com/gin-gonic/gin"
)

var (
	// 全局数据库连接
	gormDB *db.GormDB
)

// InitHandlers 初始化处理函数
func InitHandlers(db *db.GormDB) {
	gormDB = db
}

// HealthCheck 健康检查处理函数
func HealthCheck(c *gin.Context) {
	if err := gormDB.Ping(); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// Ping 测试连接处理函数
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
