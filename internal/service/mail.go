package service

import (
	"github.com/LucienVen/tech-backend/internal/db"
	"github.com/LucienVen/tech-backend/internal/repository"
)

type MailService struct {
	redis    *db.RedisClient
	userRepo repository.UserRepository
}
