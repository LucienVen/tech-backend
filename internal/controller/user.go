package controller

import (
	"net/http"

	"github.com/LucienVen/tech-backend/internal/db"
	"github.com/LucienVen/tech-backend/internal/entity"
	"github.com/LucienVen/tech-backend/internal/errors"
	"github.com/LucienVen/tech-backend/internal/request"
	"github.com/LucienVen/tech-backend/internal/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 检查用户名或邮箱是否已存在
	var count int64
	db.DB.Model(&entity.User{}).
		Where("username = ? OR email = ?", req.Username, req.Email).
		Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名或邮箱已存在"})
		return
	}

	// 加密密码
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}

	// 创建用户
	user := entity.NewUser(req.Username, req.Username, string(hash), "", req.Email)
	user.Status = entity.UserStatusInactive // 初始为未激活

	if err := db.DB.Create(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func (c *UserController) register(ctx *gin.Context) {

}
