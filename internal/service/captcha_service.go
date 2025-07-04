package service

import (
	"context"
	"time"

	"github.com/LucienVen/tech-backend/internal/db"
	"github.com/mojocn/base64Captcha"
)

const captchaExpire = 5 * time.Minute

type CaptchaService struct {
	redis *db.RedisClient
	store base64Captcha.Store // 可选，测试时注入
}

func NewCaptchaService(redis *db.RedisClient) *CaptchaService {
	return &CaptchaService{redis: redis}
}

func (s *CaptchaService) getStore() base64Captcha.Store {
	if s.store != nil {
		return s.store
	}
	return &redisStore{redis: s.redis}
}

// Verify 实现 CaptchaVerifier 接口
func (s *CaptchaService) Verify(id, code string, clear bool) bool {
	return s.getStore().Verify(id, code, clear)
}

// Generate 生成验证码
func (s *CaptchaService) Generate() (id string, b64s string, err error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, s.getStore())
	id, b64s, _, err = captcha.Generate()
	return
}

// redisStore 实现 base64Captcha.Store
// 统一调用 RedisClient 的通用方法

type redisStore struct {
	redis *db.RedisClient
}

func (s *redisStore) Set(id, value string) error {
	return s.redis.Set(context.Background(), id, value, captchaExpire)
}
func (s *redisStore) Get(id string, clear bool) string {
	val, _ := s.redis.Get(context.Background(), id)
	if clear {
		s.redis.Del(context.Background(), id)
	}
	return val
}
func (s *redisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
