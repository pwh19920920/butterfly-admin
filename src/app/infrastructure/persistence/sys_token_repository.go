package persistence

import (
	"github.com/pwh19920920/butterfly-admin/src/app/common"
	"github.com/pwh19920920/butterfly-admin/src/app/domain/entity"
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

func (tokenRepository *SysTokenRepositoryImpl) Delete(subject string) error {
	return tokenRepository.db.Model(&entity.SysToken{}).
		Where(&entity.SysToken{Subject: subject}).
		Updates(&entity.SysToken{BaseEntity: common.BaseEntity{Deleted: 1}}).Error
}

func (tokenRepository *SysTokenRepositoryImpl) GetBySubject(subject string) (*entity.SysToken, error) {
	var token entity.SysToken
	err := tokenRepository.db.Model(&entity.SysToken{}).Where(&entity.SysToken{Subject: subject}).Find(&token).Error
	return &token, err
}
