package repository

import (
	"github.com/pwh19920920/butterfly-admin/domain/entity"
)

type SysTokenRepository interface {
	// Save 保存
	Save(token entity.SysToken) error

	// Modify 更新
	Modify(token entity.SysToken) error

	// Delete 删除
	Delete(subject string) error

	// GetBySubject 查询
	GetBySubject(subject string) (*entity.SysToken, error)
}
