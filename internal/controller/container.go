package controller

import (
	appcontext "github.com/LucienVen/tech-backend/internal/appcontext"
	"github.com/LucienVen/tech-backend/internal/repository"
	"github.com/LucienVen/tech-backend/internal/service"
)

// Container 控制器容器
type Container struct {
	Health *HealthController
	// 在这里添加其他控制器
	User    *UserController
	Captcha *CaptchaController
}

// NewContainer 创建控制器容器
func NewContainer(appCtx *appcontext.AppContext) *Container {
	captchaSvc := service.NewCaptchaService(appCtx.Redis)
	return &Container{
		Health:  NewHealthController(appCtx.DB),
		User:    NewUserController(repository.NewUserRepository(appCtx.DB), captchaSvc),
		Captcha: NewCaptchaController(captchaSvc),
	}
}
