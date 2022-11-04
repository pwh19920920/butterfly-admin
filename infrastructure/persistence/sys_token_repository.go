package persistence

import (
	"errors"
	"github.com/pwh19920920/butterfly-admin/common"
	"github.com/pwh19920920/butterfly-admin/domain/entity"
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

func (tokenRepository *SysTokenRepositoryImpl) Modify(token entity.SysToken) error {
	return tokenRepository.db.Model(&entity.SysToken{}).Updates(&token).Error
}

func (tokenRepository *SysTokenRepositoryImpl) Delete(subject string) error {
	return tokenRepository.db.Model(&entity.SysToken{}).
		Where(&entity.SysToken{Subject: subject}).
		Updates(&entity.SysToken{BaseEntity: common.BaseEntity{Deleted: common.DeletedTrue}}).Error
}

func (tokenRepository *SysTokenRepositoryImpl) GetBySubject(subject string) (*entity.SysToken, error) {
	var token entity.SysToken
	err := tokenRepository.db.Model(&entity.SysToken{}).Where(&entity.SysToken{Subject: subject}).Last(&token).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &token, err
}
