package entity

import "butterfly-admin/src/app/common"

type SysPermission struct {
	common.BaseEntity

	RoleId uint64 `json:"roleId" gorm:"column:role_id"` // 角色
	MenuId uint64 `json:"menuId" gorm:"column:menu_id"` // 菜单
	Option string `json:"option" gorm:"column:option"`  // 操作
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysPermission) TableName() string {
	return "t_sys_permission"
}
