package persistence

import (
	"butterfly-admin/src/app/domain/entity"
	"gorm.io/gorm"
)

type SysTokenRepositoryImpl struct {
	db *gorm.DB
}

func NewSysTokenRepositoryImpl(db *gorm.DB) *SysTokenRepositoryImpl {
	return &SysTokenRepositoryImpl{db: db}
}

func (tokenRepository *SysTokenRepositoryImpl) Save(token entity.SysToken) error {
	return tokenRepository.db.Model(&entity.SysToken{}).Create(&token).Error
}

func (tokenRepository *SysTokenRepositoryImpl) Delete(relationId string) error {
	return tokenRepository.db.Model(&entity.SysToken{}).Updates(&entity.SysToken{Deleted: 1}).Error
}

func (tokenRepository *SysTokenRepositoryImpl) GetByRelationId(relationId string) (*entity.SysToken, error) {
	var token entity.SysToken
	err := tokenRepository.db.Model(&entity.SysToken{}).Where(&entity.SysToken{RelationId: relationId}).Find(&token).Error
	return &token, err
}
