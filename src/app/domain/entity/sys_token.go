package entity

import (
	"butterfly-admin/src/app/common"
	"encoding/json"
)

type SysToken struct {
	common.BaseEntity

	Secret     string `json:"secret" gorm:"column:secret"`          // 密钥
	UserId     uint64 `json:"userId" gorm:"column:user_id"`         // 用户
	RelationId string `json:"relationId" gorm:"column:relation_id"` // 关联id
	Deleted    int    `json:"deleted" gorm:"column:deleted"`        // 删除标记
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysToken) TableName() string {
	return "t_sys_token"
}

// 序列化
func (t SysToken) Marshal() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// 反序列化
func (SysToken) UnMarshal(text string) *SysToken {
	var token SysToken
	_ = json.Unmarshal([]byte(text), &token)
	return &token
}
