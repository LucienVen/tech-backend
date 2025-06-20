package service

import "github.com/LucienVen/tech-backend/internal/entity"

// 用户状态事件常量
const (
	UserEventActivate = "activate" // 激活
	UserEventDisable  = "disable"  // 封号
	UserEventLogout   = "logout"   // 注销
	UserEventDelete   = "delete"   // 删除
)

// 用户状态转移表
var userStatusTransitions = map[int8]map[string]int8{
	entity.UserStatusInactive: {
		UserEventActivate: entity.UserStatusActive, // 激活
	},
	entity.UserStatusActive: {
		UserEventDisable: entity.UserStatusDisabled, // 封号
		UserEventLogout:  entity.UserStatusLogout,   // 注销
		UserEventDelete:  entity.UserStatusDeleted,  // 删除
	},
	entity.UserStatusDisabled: {
		UserEventActivate: entity.UserStatusActive,  // 解封
		UserEventDelete:   entity.UserStatusDeleted, // 删除
	},
	entity.UserStatusLogout: {
		UserEventActivate: entity.UserStatusActive,  // 恢复
		UserEventDelete:   entity.UserStatusDeleted, // 删除
	},
}

// TransitionUserStatus 用户状态转移
// event: 事件名，建议使用 UserEventActivate/UserEventDisable/UserEventLogout/UserEventDelete
// 返回 true 表示转移成功，false 表示非法转移
func TransitionUserStatus(u *entity.User, event string) bool {
	if next, ok := userStatusTransitions[u.Status][event]; ok {
		u.Status = next
		return true
	}
	return false
}
