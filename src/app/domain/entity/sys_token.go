package entity

import (
	"butterfly-admin/src/app/common"
	"encoding/json"
)

type SysToken struct {
	common.BaseEntity

	Secret  string `json:"secret" gorm:"column:secret"`         // 密钥
	UserId  int64  `json:"userId,string" gorm:"column:user_id"` // 用户
	Subject string `json:"subject" gorm:"column:subject"`       // subject
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysToken) TableName() string {
	return "t_sys_token"
}

// Marshal 序列化
func (t SysToken) Marshal() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// UnMarshal 反序列化
func (SysToken) UnMarshal(text string) (*SysToken, error) {
	var token SysToken
	err := json.Unmarshal([]byte(text), &token)
	return &token, err
}
