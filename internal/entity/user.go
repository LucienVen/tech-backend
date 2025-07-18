package entity

// UserStatusInactive 表示用户未激活。
// UserStatusActive 表示用户正常。
// UserStatusDisabled 表示用户被禁用。
// UserStatusLogout 表示用户已注销。
// UserStatusDeleted 表示用户已删除。
const (
	UserStatusInactive int8 = 0 // 未激活
	UserStatusActive   int8 = 1 // 正常
	UserStatusDisabled int8 = 2 // 禁用
	UserStatusLogout   int8 = 3 // 注销
	UserStatusDeleted  int8 = 9 // 已删除
)

// User 用户实体
type User struct {
	ID       string `json:"id" gorm:"column:id"`               // 教职人员ID
	Username string `json:"username" gorm:"column:username"`   // 用户名
	NickName string `json:"nick_name" gorm:"column:nick_name"` // 昵称
	Passwd   string `json:"-" gorm:"column:passwd"`            // 密码（已加密）
	Phone    string `json:"phone" gorm:"column:phone"`         // 联系方式
	Email    string `json:"email" gorm:"column:email"`         // 邮箱地址
	Status   int8   `json:"status" gorm:"column:status"`       // '用户状态：0-未激活，1-正常，2-禁用，3-注销，9-已删除';
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
