package entity

// Grade 学生成绩表
type Grade struct {
	ID        string  `json:"id" gorm:"column:id"`                 // 成绩ID
	StudentID string  `json:"student_id" gorm:"column:student_id"` // 学生ID
	SubjectID string  `json:"subject_id" gorm:"column:subject_id"` // 学科ID
	Year      int64   `json:"year" gorm:"column:year"`             // 考试年份
	Score     float64 `json:"score" gorm:"column:score"`           // 成绩
	Term      int64   `json:"term" gorm:"column:term"`             // 学期(1-上学期,2-下学期)
	ExamID    string  `json:"exam_id" gorm:"column:exam_id"`       // 试卷ID
	BaseModel
}

// TableName 指定表名
func (Grade) TableName() string {
	return GradeTable
}
