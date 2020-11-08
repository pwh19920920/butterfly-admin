package repository

import "butterfly-admin/src/app/domain/entity"

type TokenRepository interface {
	Save(token entity.Token) error
	Delete(relationId string) error
}
