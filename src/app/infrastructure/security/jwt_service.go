package security

import (
	"butterfly-admin/src/app/config"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pwh19920920/butterfly/helper"
	"strings"
	"time"
)

type JwtServiceImpl struct {
	authConfig *config.AuthConfig
}

func NewJwtServiceImpl(authConfig *config.AuthConfig) *JwtServiceImpl {
	return &JwtServiceImpl{authConfig}
}

// 生成令牌
func (jwtService *JwtServiceImpl) GenericToken(secret, subject string) (string, error) {
	jwtSecret := []byte(secret)

	claims := jwt.StandardClaims{
		// 发布日期
		IssuedAt: time.Now().Unix(),

		// 过期时间
		ExpiresAt: time.Now().Add(time.Duration(jwtService.authConfig.ExpireTime) * time.Second).Unix(),

		// Subject
		Subject: subject,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return tokenClaims.SignedString(jwtSecret)
}

// 获取Subject
func (jwtService *JwtServiceImpl) GetSubjectFromToken(token string) (string, error) {
	// 检查数据
	typeIndex := strings.Index(token, fmt.Sprintf("%s ", jwtService.authConfig.HeaderType))
	if typeIndex != 0 {
		return "", errors.New("token数据不正确")
	}

	// 基本数据判断
	token = helper.StringHelper.SubString(token, typeIndex+1, len(token))
	startIndexDot := strings.Index(token, ".")
	lastIndexDot := strings.LastIndex(token, ".")
	if startIndexDot == -1 || lastIndexDot == -1 || startIndexDot != lastIndexDot {
		return "", errors.New("token数据不正确")
	}

	// base64数据解析
	base64Str := helper.StringHelper.SubString(token, startIndexDot+1, lastIndexDot)
	decoded, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", err
	}

	// 数据转json
	var claims jwt.StandardClaims
	err = json.Unmarshal(decoded, &claims)
	if err != nil {
		return "", err
	}

	// 获取核心数据
	return claims.Subject, nil
}
