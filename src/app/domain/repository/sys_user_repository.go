package repository

import "butterfly-admin/src/app/domain/entity"

type SysUserRepository interface {
	// 通过用户名获取用户
	GetByUsername(username string) (user *entity.SysUser)

	// 通过id获取用户
	GetById(id uint64) (user *entity.SysUser)
}
