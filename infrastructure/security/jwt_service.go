package security

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/pwh19920920/butterfly/helper"
	"strings"
	"time"
)

type JwtServiceImpl struct {
}

func NewJwtServiceImpl() *JwtServiceImpl {
	return &JwtServiceImpl{}
}

// GenericToken 生成令牌
func (jwtService *JwtServiceImpl) GenericToken(secret, subject string, expireTime time.Time) (string, error) {
	jwtSecret := []byte(secret)

	claims := jwt.StandardClaims{
		// 发布日期
		IssuedAt: time.Now().Unix(),

		// 过期时间
		ExpiresAt: expireTime.Unix(),

		// Subject
		Subject: subject,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// GetSubjectFromToken 获取Subject
func (jwtService *JwtServiceImpl) GetSubjectFromToken(token string) (string, error) {
	startIndexDot := strings.Index(token, ".")
	lastIndexDot := strings.LastIndex(token, ".")
	if startIndexDot == -1 || lastIndexDot == -1 || startIndexDot == lastIndexDot {
		return "", errors.New("token数据不正确")
	}

	// base64数据解析
	base64Str := helper.StringHelper.SubString(token, startIndexDot+1, lastIndexDot)
	decoded, err := base64.RawStdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}

	// 数据转json
	var claims = jwt.StandardClaims{}
	err = json.Unmarshal(decoded, &claims)
	if err != nil {
		return "", err
	}

	// 获取核心数据
	return claims.Subject, nil
}

// CheckToken 校验令牌
func (jwtService *JwtServiceImpl) CheckToken(token, secret string) bool {
	jwtSecret := []byte(secret)
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return err == nil && tokenClaims != nil && tokenClaims.Valid
}
