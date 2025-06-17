package entity

// Role 角色表
type Role struct {
	ID   string `json:"id" gorm:"column:id"`     // 角色ID
	Name string `json:"name" gorm:"column:name"` // 角色名
	BaseModel
}

// TableName 指定表名
func (Role) TableName() string {
	return RoleTable
}
