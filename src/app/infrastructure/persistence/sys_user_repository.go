package persistence

import (
	"butterfly-admin/src/app/domain/entity"
	"gorm.io/gorm"
)

type SysUserRepositoryImpl struct {
	db *gorm.DB
}

func NewSysUserRepositoryImpl(db *gorm.DB) *SysUserRepositoryImpl {
	return &SysUserRepositoryImpl{db: db}
}

func (userRepository *SysUserRepositoryImpl) GetByUsername(username string) *entity.SysUser {
	var user entity.SysUser
	userRepository.db.Model(&entity.SysUser{}).Where("username = ?", username).First(&user)
	return &user
}

func (userRepository *SysUserRepositoryImpl) GetById(id uint64) *entity.SysUser {
	var user entity.SysUser
	userRepository.db.Model(&entity.SysUser{}).Where("id = ?", id).First(&user)
	return &user
}
