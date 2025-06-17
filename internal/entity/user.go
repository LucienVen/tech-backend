package entity

// User 用户实体
type User struct {
	ID       string `json:"id" gorm:"column:id"`               // 教职人员ID
	Username string `json:"username" gorm:"column:username"`   // 用户名
	NickName string `json:"nick_name" gorm:"column:nick_name"` // 昵称
	Passwd   string `json:"-" gorm:"column:passwd"`            // 密码（已加密）
	Phone    string `json:"phone" gorm:"column:phone"`         // 联系方式
	Email    string `json:"email" gorm:"column:email"`         // 邮箱地址
	BaseModel
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// NewUser 创建新用户
func NewUser(username, nickName, passwd, phone, email string) *User {
	return &User{
		BaseModel: GenBaseModel("system", "system"),
		Username:  username,
		NickName:  nickName,
		Passwd:    passwd,
		Phone:     phone,
		Email:     email,
	}
}
