package security

import (
	"butterfly-admin/src/app/config/auth"
)

type TokenService interface {
	// GenericToken 生成令牌
	GenericToken(authConfig *auth.Config, secret, subject string) (string, error)

	// GetSubjectFromToken 获取Subject
	GetSubjectFromToken(token string) (string, error)

	// CheckToken 校验令牌
	CheckToken(token, secret string) bool
}
