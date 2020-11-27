package repository

import "butterfly-admin/src/app/domain/entity"

type SysTokenRepository interface {
	// 保存
	Save(token entity.SysToken) error

	// 删除
	Delete(subject string) error

	// 查询
	GetBySubject(subject string) (*entity.SysToken, error)
}
