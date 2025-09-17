package service

import (
	"errors"
	"github.com/LucienVen/tech-backend/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (s *UserService) CheckUserExists(username, email string) (bool, error) {
	if username != "" && email != "" {
		existsUser, err1 := s.repo.ExistsByUsername(username)
		existsEmail, err2 := s.repo.ExistsByEmail(email)
		return existsUser || existsEmail, errors.Join(err1, err2)
	} else if username != "" {
		return s.repo.ExistsByUsername(username)
	} else if email != "" {
		return s.repo.ExistsByEmail(email)
	}
	return false, nil
}
