package entity

// Subject 学科信息表
type Subject struct {
	ID          string `json:"id" gorm:"column:id"`                   // 学科ID
	Name        string `json:"name" gorm:"column:name"`               // 学科名称
	Description string `json:"description" gorm:"column:description"` // 学科介绍
	DirectorID  string `json:"director_id" gorm:"column:director_id"` // 主任ID
	BaseModel
}

// TableName 指定表名
func (Subject) TableName() string {
	return SubjectTable
}
