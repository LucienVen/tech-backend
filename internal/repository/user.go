package repository

import (
	"github.com/LucienVen/tech-backend/internal/entity"
	"gorm.io/gorm"
)

// DBType 数据库类型
// 可根据实际项目放到 db 包
// type DBType string
// const (
//
//	DBTypeMySQL = "mysql"
//	DBTypePG    = "pg"
//
// )
// DBTypeMySQL 表示 MySQL 数据库类型
// DBTypePG 表示 PostgreSQL 数据库类型
type DBType string

const (
	DBTypeMySQL = "mysql"
	DBTypePG    = "pg"
)

// UserRepository 用户仓储接口
type UserRepository interface {
	Exists(username, email string) (bool, error)
	// 可扩展更多方法
}

// MySQL 实现
type userRepositoryMySQL struct {
	db *gorm.DB
}

func NewUserRepositoryMySQL(db *gorm.DB) UserRepository {
	return &userRepositoryMySQL{db: db}
}

func (r *userRepositoryMySQL) Exists(username, email string) (bool, error) {
	var count int64
	err := r.db.Model(&entity.User{}).
		Where("username = ? OR email = ?", username, email).
		Count(&count).Error
	return count > 0, err
}

// PostgreSQL 实现
type userRepositoryPG struct {
	db *gorm.DB
}

func NewUserRepositoryPG(db *gorm.DB) UserRepository {
	return &userRepositoryPG{db: db}
}

func (r *userRepositoryPG) Exists(username, email string) (bool, error) {
	var count int64
	err := r.db.Model(&entity.User{}).
		Where("username = $1 OR email = $2", username, email).
		Count(&count).Error
	return count > 0, err
}

// 工厂方法，根据类型创建
func NewUserRepository(dbType DBType, db *gorm.DB) UserRepository {
	switch dbType {
	case DBTypePG:
		return NewUserRepositoryPG(db)
	default:
		return NewUserRepositoryMySQL(db)
	}
}
