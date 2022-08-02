package security

import (
	"github.com/pwh19920920/butterfly-admin/config/auth"
	"time"
)

type TokenService interface {
	// GenericToken 生成令牌
	GenericToken(authConfig *auth.Config, secret, subject string, expireTime time.Time) (string, error)

	// GetSubjectFromToken 获取Subject
	GetSubjectFromToken(token string) (string, error)

	// CheckToken 校验令牌
	CheckToken(token, secret string) bool
}
