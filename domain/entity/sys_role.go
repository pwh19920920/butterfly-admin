package entity

import (
	"github.com/pwh19920920/butterfly-admin/common"
)

type SysRole struct {
	common.BaseEntity

	Name string `json:"name" gorm:"column:name"` // 角色名称
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysRole) TableName() string {
	return "t_sys_role"
}
