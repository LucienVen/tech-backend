package controller

import (
	"github.com/LucienVen/tech-backend/internal/db"
	"github.com/LucienVen/tech-backend/internal/errors"
	"github.com/LucienVen/tech-backend/internal/response"
	"github.com/gin-gonic/gin"
)

// UserController 用户控制器
type UserController struct {
	db *db.GormDB
}

// NewUserController 创建用户控制器实例
func NewUserController(db *db.GormDB) *UserController {
	return &UserController{
		db: db,
	}
}

// Get 获取用户
func (c *UserController) Get(ctx *gin.Context) {
	response.Success(ctx, gin.H{
		"status": "ok",
		"code":   errors.ErrCodeSuccess,
	})
	return
}
