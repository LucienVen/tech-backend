package entity

// Teacher 教职人员表
type Teacher struct {
	ID        string `json:"id" gorm:"column:id"`                 // 教职人员ID
	Name      string `json:"name" gorm:"column:name"`             // 姓名
	Age       int64  `json:"age" gorm:"column:age"`               // 年龄
	Gender    int64  `json:"gender" gorm:"column:gender"`         // 性别(1-男,2-女)
	SubjectID string `json:"subject_id" gorm:"column:subject_id"` // 所教学科ID
	Phone     string `json:"phone" gorm:"column:phone"`           // 联系方式
	Email     string `json:"email" gorm:"column:email"`           // 邮箱地址
	Passwd    string `json:"passwd" gorm:"column:passwd"`         // 密码（已加密）
	Level     int64  `json:"level" gorm:"column:level"`           // 负责年级
	BaseModel
}

// TableName 指定表名
func (Teacher) TableName() string {
	return TeacherTable
}
