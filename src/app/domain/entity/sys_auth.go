package entity

import "butterfly-admin/src/app/common"

type SysAuth struct {
	common.BaseEntity

	UserId  int64 `json:"userId,string" gorm:"column:user_id"`
	RoleId  int64 `json:"roleId,string" gorm:"column:role_id"`
	Deleted int   `json:"deleted" gorm:"column:deleted"` // 删除标记
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysAuth) TableName() string {
	return "t_sys_user_role"
}
