package repository

import (
	"github.com/LucienVen/tech-backend/internal/db"
	db2 "github.com/LucienVen/tech-backend/internal/db"
	"github.com/LucienVen/tech-backend/internal/entity"
	"gorm.io/gorm"
)

// UserRepository 用户仓储接口
type UserRepository interface {
	Exists(username, email string) (bool, error)
	Create(user *entity.User) error
	// 可扩展更多方法
}

// MySQL 实现
type userRepositoryMySQL struct {
	db *gorm.DB
}

// NewUserRepositoryMySQL 创建 MySQL 用户仓储
func NewUserRepositoryMySQL(db *gorm.DB) UserRepository {
	return &userRepositoryMySQL{db: db}
}

func (r *userRepositoryMySQL) Exists(username, email string) (bool, error) {
	var count int64
	db := r.db.Model(&entity.User{})
	if username != "" && email != "" {
		db = db.Where("username = ? OR email = ?", username, email)
	} else if username != "" {
		db = db.Where("username = ?", username)
	} else if email != "" {
		db = db.Where("email = ?", email)
	} else {
		return false, nil // 两个都空，直接返回不存在
	}
	err := db.Count(&count).Error
	return count > 0, err
}

func (r *userRepositoryMySQL) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// PostgreSQL 实现
type userRepositoryPG struct {
	db *gorm.DB
}

// NewUserRepositoryPG 创建 PostgreSQL 用户仓储
func NewUserRepositoryPG(db *gorm.DB) UserRepository {
	return &userRepositoryPG{db: db}
}

func (r *userRepositoryPG) Exists(username, email string) (bool, error) {
	var count int64
	db := r.db.Model(&entity.User{})
	if username != "" && email != "" {
		db = db.Where("username = $1 OR email = $2", username, email)
	} else if username != "" {
		db = db.Where("username = $1", username)
	} else if email != "" {
		db = db.Where("email = $1", email)
	} else {
		return false, nil // 两个都空，直接返回不存在
	}
	err := db.Count(&count).Error
	return count > 0, err
}

func (r *userRepositoryPG) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// NewUserRepository 工厂方法，根据 db.DB 类型创建
func NewUserRepository(db db.DB) UserRepository {
	switch db.GetType() {
	case db2.DBTypeMySQL:
		return NewUserRepositoryMySQL(db.GetDB())
	case db2.DBTypePG:
		return NewUserRepositoryPG(db.GetDB())
	default:
		panic("unsupported db type for UserRepository")
	}
}
