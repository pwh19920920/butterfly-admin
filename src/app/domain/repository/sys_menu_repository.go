package repository

import (
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/types"
)

type SysMenuRepository interface {
	// 保存
	Save(menu *entity.SysMenu) error

	// 更新
	UpdateById(id int64, menu *entity.SysMenu) error

	// 删除
	Delete(id int64) error

	// 分页查询
	Select(req *types.SysMenuQueryRequest) (int64, []entity.SysMenu, error)

	// 查询全部
	SelectAll() ([]entity.SysMenu, error)
}
