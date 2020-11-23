package entity

import "butterfly-admin/src/app/common"

type SysUser struct {
	common.BaseEntity

	Username string `json:"updatedAt" gorm:"column:username"` // 用户
	Password string `json:"updatedAt" gorm:"column:password"` // 密码
	Salt     string `json:"salt" gorm:"column:salt"`          // 密码盐
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysUser) TableName() string {
	return "t_sys_user"
}
