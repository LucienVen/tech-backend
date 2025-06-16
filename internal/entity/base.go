package entity

import "time"

// BaseModel 基础模型，包含通用字段
type BaseModel struct {
	IsDelete   int64  `json:"is_delete" gorm:"column:is_delete"`     // 是否删除（逻辑删除标记）
	Creator    string `json:"creator" gorm:"column:creator"`         // 创建者
	Updater    string `json:"updater" gorm:"column:updater"`         // 更新者
	CreateTime int64  `json:"create_time" gorm:"column:create_time"` // 创建时间
	UpdateTime int64  `json:"update_time" gorm:"column:update_time"` // 更新时间
}

// GenWhenUpdate 生成更新时的基础字段
func (bm *BaseModel) GenWhenUpdate(handler string) *BaseModel {
	bm.Updater = handler
	bm.UpdateTime = time.Now().Unix()
	return bm
}

// GenWhenUpdateMap 生成更新时的字段映射
func (bm *BaseModel) GenWhenUpdateMap(handler string) map[string]interface{} {
	bm.GenWhenUpdate(handler)
	return map[string]interface{}{
		"updater":     bm.Updater,
		"update_time": bm.UpdateTime,
	}
}

// GenBaseModel 生成基础模型
func GenBaseModel(c, u string) BaseModel {
	timenow := time.Now().Unix()
	return BaseModel{
		IsDelete:   0,
		Creator:    c,
		Updater:    u,
		CreateTime: timenow,
		UpdateTime: timenow,
	}
}
