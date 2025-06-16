package repository

import (
	"time"

	"github.com/LucienVen/tech-backend/internal/entity"
	"gorm.io/gorm"
)

// ClassRepository 班级仓储接口实现
type ClassRepository struct {
	db *gorm.DB
}

// 更新字段常量
const (
	fieldMainTeacherID = "main_teacher_id"
	fieldUpdater       = "updater"
	fieldUpdateTime    = "update_time"
)

// NewClassRepository 创建班级仓储实例
func NewClassRepository(db *gorm.DB) *ClassRepository {
	return &ClassRepository{db: db}
}

// GetAll 获取所有班级信息
func (r *ClassRepository) GetAll() ([]entity.Class, error) {
	var classes []entity.Class
	err := r.db.Where("is_delete = ?", 0).Find(&classes).Error
	return classes, err
}

// UpdateMainTeacherID 更新班级班主任
func (r *ClassRepository) UpdateMainTeacherID(classID string, teacherID string, handler string) error {
	// 构建更新数据
	class := entity.Class{
		MainTeacherID: teacherID,
		BaseModel: entity.BaseModel{
			Updater:    handler,
			UpdateTime: time.Now().Unix(),
		},
	}

	// 执行更新
	return r.db.Model(&entity.Class{}).
		Where("id = ?", classID).
		Select(fieldMainTeacherID, fieldUpdater, fieldUpdateTime).
		Updates(class).Error
}
