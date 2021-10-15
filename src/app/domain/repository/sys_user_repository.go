package repository

import (
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/types"
)

type SysUserRepository interface {
	// GetByUsername 通过用户名获取用户
	GetByUsername(username string) (*entity.SysUser, error)

	// GetById 通过id获取用户
	GetById(id int64) (*entity.SysUser, error)

	// Select 分页查询用户
	Select(request *types.SysUserQueryRequest) (int64, []entity.SysUser, error)

	// Create 创建
	Create(user *entity.SysUser) error

	// Modify 更新
	Modify(user *entity.SysUser) error
}
