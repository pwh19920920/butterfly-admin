package entity

import "butterfly-admin/src/app/common"

type SysUser struct {
	common.BaseEntity

	Name     string `json:"name" gorm:"column:name"`         // 名称
	Avatar   string `json:"avatar"  gorm:"column:avatar"`    // 头像
	Username string `json:"username" gorm:"column:username"` // 用户
	Password string `json:"password" gorm:"column:password"` // 密码
	Salt     string `json:"salt" gorm:"column:salt"`         // 密码盐
	Roles    string `json:"roles" gorm:"column:roles"`       // 角色串
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysUser) TableName() string {
	return "t_sys_user"
}
