package persistence

import (
	"butterfly-admin/src/app/domain/repository"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository  repository.SysUserRepository
	TokenRepository repository.SysTokenRepository
	db              *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:              db,
		UserRepository:  NewSysUserRepositoryImpl(db),
		TokenRepository: NewSysTokenRepositoryImpl(db),
	}
}
