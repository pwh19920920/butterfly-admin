package entity

import (
	"butterfly-admin/src/app/common"
	"encoding/json"
)

type Token struct {
	Id         uint64            `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt  *common.LocalTime `json:"createdAt" gorm:"index;column:created_at"`
	UpdatedAt  *common.LocalTime `json:"updatedAt" gorm:"column:updated_at"`
	Secret     string            `json:"secret" gorm:"column:secret"`
	UserId     uint64            `json:"userId" gorm:"column:user_id"`
	RelationId string            `json:"relationId" gorm:"column:relation_id"`
	Deleted    int               `json:"deleted" gorm:"column:deleted"`
}

// TableName 会将 User 的表名重写为 `profiles`
func (Token) TableName() string {
	return "t_sys_token"
}

// 序列化
func (t Token) Marshal() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// 反序列化
func (Token) UnMarshal(text string) *Token {
	var token Token
	_ = json.Unmarshal([]byte(text), &token)
	return &token
}
