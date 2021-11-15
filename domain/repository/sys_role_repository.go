package repository

import (
	"github.com/pwh19920920/butterfly-admin/domain/entity"
	"github.com/pwh19920920/butterfly-admin/types"
)

type SysRoleRepository interface {
	// Save 保存
	Save(permissions *[]entity.SysPermission, role *entity.SysRole) error

	// UpdateById 更新
	UpdateById(id int64, permissions *[]entity.SysPermission, role *entity.SysRole) error

	// Delete 删除
	Delete(id int64) error

	// Select 分页查询
	Select(req *types.SysRoleQueryRequest) (int64, []entity.SysRole, error)

	// SelectAll 查询全部
	SelectAll() ([]entity.SysRole, error)

	//
}
