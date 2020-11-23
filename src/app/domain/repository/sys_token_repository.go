package repository

import "butterfly-admin/src/app/domain/entity"

type SysTokenRepository interface {
	// 保存
	Save(token entity.SysToken) error

	// 删除
	Delete(relationId string) error

	// 查询
	GetByRelationId(relationId string) (*entity.SysToken, error)
}
