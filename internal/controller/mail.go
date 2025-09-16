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

func (m *MailController) VerifyEmail(ctx *gin.Context) {

}
