package repository

import "butterfly-admin/src/app/domain/entity"

type TokenRepository interface {
	// 保存
	Save(token entity.Token) error

	// 删除
	Delete(relationId string) error

	// 查询
	GetByRelationId(relationId string) (*entity.Token, error)
}
