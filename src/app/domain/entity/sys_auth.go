package entity

import "butterfly-admin/src/app/common"

type SysAuth struct {
	common.BaseEntity

	UserId  uint64 `json:"userId" gorm:"column:user_id"`
	RoleId  uint64 `json:"roleId" gorm:"column:role_id"`
	Deleted int    `json:"deleted" gorm:"column:deleted"` // 删除标记
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysAuth) TableName() string {
	return "t_sys_user_role"
}
