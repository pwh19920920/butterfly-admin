package types

import "github.com/pwh19920920/butterfly-admin/src/app/domain/entity"

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type SysMenuPermissionForUser struct {
	entity.SysUser
	Menus       []SysMenuPermissionForUserMenu `json:"menus"`
	Permissions []string                       `json:"permissions"`
	Codes       []string                       `json:"codes"`
}

type SysMenuPermissionForUserMenu struct {
	Id        int64                          `json:"id"`        // id
	Name      string                         `json:"name"`      // 菜单名称
	Path      string                         `json:"path"`      // 菜单路径
	Icon      string                         `json:"icon"`      // 菜单图标
	Component string                         `json:"component"` // 菜单组件
	Routes    []SysMenuPermissionForUserMenu `json:"routes"`
}
