package persistence

import (
	"butterfly-admin/src/app/config"
	"butterfly-admin/src/app/domain/repository"
)

type Repository struct {
	SysUserRepository       repository.SysUserRepository
	SysTokenRepository      repository.SysTokenRepository
	SysMenuRepository       repository.SysMenuRepository
	SysRoleRepository       repository.SysRoleRepository
	SysPermissionRepository repository.SysPermissionRepository
	SysMenuOptionRepository repository.SysMenuOptionRepository
}

func NewRepository(config config.Config) *Repository {
	return &Repository{
		SysMenuOptionRepository: NewSysMenuOptionRepositoryImpl(config.DatabaseForGorm),
		SysPermissionRepository: NewSysPermissionRepositoryImpl(config.DatabaseForGorm),
		SysUserRepository:       NewSysUserRepositoryImpl(config.DatabaseForGorm),
		SysTokenRepository:      NewSysTokenRepositoryImpl(config.DatabaseForGorm),
		SysMenuRepository:       NewSysMenuRepositoryImpl(config.DatabaseForGorm),
		SysRoleRepository:       NewSysRoleRepositoryImpl(config.DatabaseForGorm),
	}
}
