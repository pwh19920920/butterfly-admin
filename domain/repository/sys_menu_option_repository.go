package repository

import (
	"github.com/pwh19920920/butterfly-admin/domain/entity"
)

type SysMenuOptionRepository interface {

	// SelectByMenuId 通过菜单id搜索
	SelectByMenuId(menuId int64) ([]entity.SysMenuOption, error)

	// SelectAll 查询全部
	SelectAll() ([]entity.SysMenuOption, error)

	// SelectByIds 批量查询
	SelectByIds(ids []int64) ([]entity.SysMenuOption, error)
}
