package controller

import (
	"context"
	"time"

	"github.com/LucienVen/tech-backend/internal/db"
	"github.com/LucienVen/tech-backend/internal/errors"
	"github.com/LucienVen/tech-backend/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// RedisStore 实现 base64Captcha.Store 接口，持有 *db.RedisClient
var _ base64Captcha.Store = (*RedisStore)(nil)

type RedisStore struct {
	redis *db.RedisClient
}

func (s *RedisStore) Set(id, value string) error {
	return s.redis.GetClient().Set(context.Background(), id, value, 5*time.Minute).Err()
}
func (s *RedisStore) Get(id string, clear bool) string {
	val, _ := s.redis.GetClient().Get(context.Background(), id).Result()
	if clear {
		s.redis.GetClient().Del(context.Background(), id)
	}
	return val
}
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}

// CaptchaController 验证码控制器
type CaptchaController struct {
	store base64Captcha.Store
}

func NewCaptchaController(redisClient *db.RedisClient) *CaptchaController {
	return &CaptchaController{
		store: &RedisStore{redis: redisClient},
	}
}

// GetCaptcha 生成图片验证码
func (c *CaptchaController) GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, c.store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		response.Error(ctx, errors.ErrCodeSystemError, "验证码生成失败")
		return
	}
	response.Success(ctx, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
