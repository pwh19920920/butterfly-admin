package persistence

import (
	"butterfly-admin/src/app/domain/entity"
	"gorm.io/gorm"
)

type TokenRepo struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepo {
	return &TokenRepo{db: db}
}

func (tokenRepository *TokenRepo) Save(token entity.Token) error {
	return tokenRepository.db.Model(&entity.Token{}).Create(&token).Error
}

func (tokenRepository *TokenRepo) Delete(relationId string) error {
	return tokenRepository.db.Model(&entity.Token{}).Updates(&entity.Token{Deleted: true}).Error
}

func (tokenRepository *TokenRepo) GetByRelationId(relationId string) (*entity.Token, error) {
	var token entity.Token
	err := tokenRepository.db.Model(&entity.Token{}).Where(&entity.Token{RelationId: relationId}).Find(&token).Error
	return &token, err
}
