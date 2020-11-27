package entity

import "butterfly-admin/src/app/common"

type SysRole struct {
	common.BaseEntity

	Name    string `json:"name" gorm:"column:name"`       // 角色名称
	Deleted int    `json:"deleted" gorm:"column:deleted"` // 删除标记
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysRole) TableName() string {
	return "t_sys_role"
}
