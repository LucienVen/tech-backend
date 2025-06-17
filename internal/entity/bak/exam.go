package entity

// Exam 考试表
type Exam struct {
	ID        string `json:"id" gorm:"column:id"`                 // 考试ID
	Name      string `json:"name" gorm:"column:name"`             // 考试名称
	SubjectID string `json:"subject_id" gorm:"column:subject_id"` // 学科ID
	Year      int64  `json:"year" gorm:"column:year"`             // 考试年份
	BaseModel
}

// TableName 指定表名
func (Exam) TableName() string {
	return ExamTable
}
