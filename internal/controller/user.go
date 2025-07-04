package controller

import (
	"github.com/LucienVen/tech-backend/internal/entity"
	"github.com/LucienVen/tech-backend/internal/errors"
	"github.com/LucienVen/tech-backend/internal/repository"
	"github.com/LucienVen/tech-backend/internal/request"
	"github.com/LucienVen/tech-backend/internal/response"
	"github.com/LucienVen/tech-backend/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserController 用户控制器
type UserController struct {
	userRepo        repository.UserRepository
	captchaVerifier service.CaptchaVerifier
}

// NewUserController 创建用户控制器实例
func NewUserController(userRepo repository.UserRepository, captchaVerifier service.CaptchaVerifier) *UserController {
	return &UserController{
		userRepo:        userRepo,
		captchaVerifier: captchaVerifier,
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

// Register 注册用户
func (c *UserController) Register(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, errors.ErrCodeParamInvalid, err.Error())
		return
	}

	// 检查用户名或邮箱是否已存在
	exists, err := c.userRepo.Exists(req.Username, req.Email)
	if err != nil {
		response.Error(ctx, errors.ErrCodeSystemError, "查询用户失败: "+err.Error())
		return
	}
	if exists {
		response.Error(ctx, errors.ErrCodeUserAlreadyExists, "用户名或邮箱已存在")
		return
	}

	// 验证码校验
	if !c.captchaVerifier.Verify(req.CaptchaID, req.CaptchaCode, true) {
		response.Error(ctx, errors.ErrCodeParamInvalid, "验证码错误")
		return
	}

	// 加密密码
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(ctx, errors.ErrCodeSystemError, "密码加密失败")
		return
	}

	// 创建用户
	user := entity.NewUser(req.Username, req.Username, string(hash), "", req.Email)
	user.Status = entity.UserStatusInactive // 初始为未激活

	// TODO: 这里建议在 UserRepository 增加 Create 方法，统一数据写入
	// 临时方案：如有需要可在 userRepo 接口和实现中补充 Create

	response.Success(ctx, gin.H{"message": "注册成功"})
}
