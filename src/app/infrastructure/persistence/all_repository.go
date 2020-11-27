package persistence

import (
	"butterfly-admin/src/app/domain/repository"
	"gorm.io/gorm"
)

type Repository struct {
	SysUserRepository  repository.SysUserRepository
	SysTokenRepository repository.SysTokenRepository
	SysMenuRepository  repository.SysMenuRepository
	db                 *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:                 db,
		SysUserRepository:  NewSysUserRepositoryImpl(db),
		SysTokenRepository: NewSysTokenRepositoryImpl(db),
		SysMenuRepository:  NewSysMenuRepositoryImpl(db),
	}
}
