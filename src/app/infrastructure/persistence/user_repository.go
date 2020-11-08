package persistence

import (
	"butterfly-admin/src/app/domain/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (userRepository *UserRepo) GetUser(username string) *entity.User {
	var user entity.User
	userRepository.db.Model(&entity.User{}).Where("username = ?", username).First(&user)
	return &user
}
