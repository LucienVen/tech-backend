package entity

// Student 学生信息表
type Student struct {
	ID      string `json:"id" gorm:"column:id"`             // 学生ID
	Name    string `json:"name" gorm:"column:name"`         // 姓名
	Gender  int64  `json:"gender" gorm:"column:gender"`     // 性别
	ClassID string `json:"class_id" gorm:"column:class_id"` // 班级ID
	Phone   string `json:"phone" gorm:"column:phone"`       // 联系方式
	Email   string `json:"email" gorm:"column:email"`       // 邮箱
	Passwd  string `json:"passwd" gorm:"column:passwd"`     // 登录密码
	BaseModel
}

// TableName 指定表名
func (Student) TableName() string {
	return StudentTable
}
