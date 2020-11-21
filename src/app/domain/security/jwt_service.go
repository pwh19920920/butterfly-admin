package security

import "butterfly-admin/src/app/config"

type JwtService interface {
	// 生成令牌
	GenericToken(authConfig *config.AuthConfig, secret, subject string) (string, error)

	// 获取Subject
	GetSubjectFromToken(token string) (string, error)

	// 校验令牌
	CheckToken(token, secret string) bool
}
