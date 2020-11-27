package entity

import "butterfly-admin/src/app/common"

type SysMenu struct {
	common.BaseEntity

	Name      string `json:"name" gorm:"column:name"`            // 菜单名称
	Path      string `json:"path" gorm:"column:path"`            // 菜单路径
	Icon      string `json:"icon" gorm:"column:icon"`            // 菜单图标
	Component string `json:"component" gorm:"column:component"`  // 菜单组件
	Sort      string `json:"sort" gorm:"column:sort"`            // 菜单排序
	Option    string `json:"option" gorm:"column:option"`        // 菜单操作
	Parent    int64  `json:"parent,string" gorm:"column:parent"` // 上级目录
	Deleted   int    `json:"deleted" gorm:"column:deleted"`      // 删除标记
}

// TableName 会将 User 的表名重写为 `profiles`
func (SysMenu) TableName() string {
	return "t_sys_menu"
}
