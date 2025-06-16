package entity

// Class 班级表
type Class struct {
	ID            string `json:"id" gorm:"column:id"`                           // 班级ID
	ClassNum      int64  `json:"class_num" gorm:"column:class_num"`             // 班级序号
	GradeLevel    int64  `json:"grade_level" gorm:"column:grade_level"`         // 年级
	MainTeacherID string `json:"main_teacher_id" gorm:"column:main_teacher_id"` // 班主任
	BaseModel
}

// TableName 指定表名
func (Class) TableName() string {
	return ClassTable
}
