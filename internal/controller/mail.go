package controller

import (
	"github.com/LucienVen/tech-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type MailController struct {
	captchaService service.CaptchaVerifier
}

func NewMailController(captchaService service.CaptchaVerifier) *MailController {
	return &MailController{
		captchaService: captchaService,
	}
}

// 检验
func (m *MailController) VerifyEmail(ctx *gin.Context) {
	//token := ctx.Query("token")

}

// 发送验证邮件
func (m *MailController) SendVerifyEmail(ctx *gin.Context) {

}
