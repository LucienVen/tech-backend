package controller

import (
	"github.com/LucienVen/tech-backend/internal/errors"
	"github.com/LucienVen/tech-backend/internal/response"
	"github.com/LucienVen/tech-backend/internal/service"
	"github.com/gin-gonic/gin"
)

// CaptchaController 负责验证码相关接口
// 依赖 service.CaptchaVerifier 进行业务处理
// 只处理 HTTP 路由与参数转发
type CaptchaController struct {
	captchaService service.CaptchaVerifier
}

// NewCaptchaController 创建 CaptchaController 实例
func NewCaptchaController(captchaService service.CaptchaVerifier) *CaptchaController {
	return &CaptchaController{
		captchaService: captchaService,
	}
}

// GetCaptcha 生成图片验证码并返回 base64 图片字符串
// @route GET /captcha
func (c *CaptchaController) GetCaptcha(ctx *gin.Context) {
	id, b64s, err := c.captchaService.Generate()
	if err != nil {
		response.Error(ctx, errors.ErrCodeSystemError, "验证码生成失败")
		return
	}
	response.Success(ctx, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

// Verify 校验验证码
// 仅供内部调用，不暴露为 HTTP 接口
func (c *CaptchaController) Verify(id, code string, clear bool) bool {
	return c.captchaService.Verify(id, code, clear)
}
