package repository

import "butterfly-admin/src/app/domain/entity"

type UserRepository interface {
	GetUser(username string) (user *entity.User)
}
