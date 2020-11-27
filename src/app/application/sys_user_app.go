package application

import (
	"butterfly-admin/src/app/config"
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/domain/security"
	"butterfly-admin/src/app/infrastructure/persistence"
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/pwh19920920/butterfly/helper"
	"github.com/sirupsen/logrus"
	"strings"
)

// 忽略的地址
var ignorePathMap *map[string]bool

type SysUserApplication struct {
	sequence       *snowflake.Node
	repository     *persistence.Repository
	encoderService security.EncodeService
	jwtService     security.JwtService
	authConfig     *config.AuthConfig
}

// 退出
func (l *SysUserApplication) Logout(subject string) error {
	return l.repository.SysTokenRepository.Delete(subject)
}

// 登陆
func (l *SysUserApplication) Login(username, password string) (ticket string, err error) {
	user := l.repository.SysUserRepository.GetByUsername(username)
	if user == nil {
		return "", errors.New("用户不存在")
	}

	// 检查密码
	encPassword := l.encoderService.Encode(password, user.Salt)
	if encPassword != user.Password {
		return "", errors.New("用户密码不正确")
	}

	// 生成令牌数据
	return l.genericToken(user.Id)
}

// 获取配置名称
func (l *SysUserApplication) GetHeaderName() string {
	return l.authConfig.HeaderName
}

// 检查并获取用户id
func (l *SysUserApplication) CheckAndGetTicket(token string) (*entity.SysToken, error) {
	// 取出票据id
	token, err := l.parseToken(token)
	if err != nil {
		return nil, errors.New("token数据不正确")
	}

	subject, err := l.jwtService.GetSubjectFromToken(token)
	if err != nil {
		return nil, err
	}

	// 取出票据对象
	ticket, err := l.repository.SysTokenRepository.GetBySubject(subject)
	if err != nil {
		return nil, err
	}

	// 判断票据是否为空， 并校验
	if ticket == nil {
		return nil, errors.New("token不存在")
	}

	if !l.jwtService.CheckToken(token, ticket.Secret) {
		return nil, errors.New("令牌校验失败")
	}

	// 校验成功，返回用户id
	return ticket, nil
}

// 刷新令牌
func (l *SysUserApplication) RefreshToken(userId int64, subject, token string) (string, error) {
	// 取出票据id
	token, err := l.parseToken(token)
	if err != nil {
		return "", errors.New("token数据不正确")
	}

	// 生成令牌数据
	return l.genericToken(userId)
}

// 获取用户拥有的权限信息
func (l *SysUserApplication) GetUserPermission(userId int64) *map[string]bool {
	specMap := make(map[string]bool)
	return &specMap
}

// 获取忽略auth的地址
func (l *SysUserApplication) GetIgnorePaths() *map[string]bool {
	if ignorePathMap == nil {
		dataMap := make(map[string]bool)
		for _, v := range l.authConfig.IgnorePath {
			dataMap[v] = true
			ignorePathMap = &dataMap
		}
	}
	return ignorePathMap
}

// 生成令牌
func (l *SysUserApplication) genericToken(userId int64) (string, error) {
	// 生成保存密钥
	secret := uuid.New().String()
	subject := uuid.New().String()

	// 保存用户信息与令牌之间的关系
	// subject -> userId
	// subject -> secret
	// userId -> subject
	err := l.repository.SysTokenRepository.Save(entity.SysToken{
		Secret:  secret,
		Subject: subject,
		UserId:  userId,
	})

	// 判定是否保存失败
	if err != nil {
		logrus.Error(err)
		return "", errors.New("密钥保存失败")
	}

	// 生成令牌数据
	return l.jwtService.GenericToken(l.authConfig, secret, subject)
}

// 从header中解析令牌
func (l *SysUserApplication) parseToken(token string) (string, error) {
	// 检查数据
	typeKey := fmt.Sprintf("%s ", l.authConfig.HeaderType)
	typeIndex := strings.Index(token, typeKey)
	if typeIndex != 0 {
		return "", errors.New("token数据不正确")
	}

	// 取出票据id
	return helper.StringHelper.SubString(token, len(typeKey), len(token)), nil
}
