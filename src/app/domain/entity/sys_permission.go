package entity

import "butterfly-admin/src/app/common"

type SysPermission struct {
	common.BaseEntity

	RoleId  int64  `json:"roleId,string" gorm:"column:role_id"` // 角色
	MenuId  int64  `json:"menuId,string" gorm:"column:menu_id"` // 菜单
	Option  string `json:"option" gorm:"column:option"`         // 操作
	Deleted int    `json:"deleted" gorm:"column:deleted"`       // 删除标记
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysPermission) TableName() string {
	return "t_sys_permission"
}
