package persistence

import (
	"butterfly-admin/src/app/domain/repository"
	"gorm.io/gorm"
)

type Repository struct {
	UserRepository  repository.UserRepository
	TokenRepository repository.TokenRepository
	db              *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:              db,
		UserRepository:  NewUserRepository(db),
		TokenRepository: NewTokenRepository(db),
	}
}
