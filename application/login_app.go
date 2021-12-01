package application

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/pwh19920920/butterfly-admin/config/auth"
	"github.com/pwh19920920/butterfly-admin/domain/entity"
	"github.com/pwh19920920/butterfly-admin/domain/security"
	"github.com/pwh19920920/butterfly-admin/infrastructure/persistence"
	"github.com/pwh19920920/butterfly-admin/types"
	"github.com/pwh19920920/butterfly/helper"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

// 忽略的地址
var ignorePathMap = make(map[string]bool, 0)
var ignorePrefixPaths = make([]string, 0)
var commonPathMap = make(map[string]bool, 0)
var initPath = false

func init() {
	ignorePathMap["POST - /api/login"] = true
	commonPathMap["POST - /api/logout"] = true
	commonPathMap["POST - /api/refresh"] = true
	commonPathMap["GET - /api/currentUser"] = true
}

type LoginApplication struct {
	sequence       *snowflake.Node
	repository     *persistence.Repository
	encoderService security.EncodeService
	tokenService   security.TokenService
	authConfig     *auth.Config
}

// Logout 退出
func (application *LoginApplication) Logout(subject string) error {
	return application.repository.SysTokenRepository.Delete(subject)
}

// Login 登陆
func (application *LoginApplication) Login(username, password string) (ticket string, err error) {
	user, err := application.repository.SysUserRepository.GetByUsername(username)
	if err != nil || user == nil {
		return "", errors.New("用户不存在或者获取失败")
	}

	// 检查密码
	encPassword := application.encoderService.Encode(password, user.Salt)
	if encPassword != user.Password {
		return "", errors.New("用户密码不正确")
	}

	// 生成令牌数据
	return application.genericToken(user.Id)
}

// GetHeaderName 获取配置名称
func (application *LoginApplication) GetHeaderName() string {
	return application.authConfig.HeaderName
}

// CheckAndGetTicket 检查并获取用户id
func (application *LoginApplication) CheckAndGetTicket(token string) (*entity.SysToken, error) {
	// 取出票据id
	token, err := application.parseToken(token)
	if err != nil {
		return nil, errors.New("token数据不正确")
	}

	subject, err := application.tokenService.GetSubjectFromToken(token)
	if err != nil {
		return nil, err
	}

	// 取出票据对象
	ticket, err := application.repository.SysTokenRepository.GetBySubject(subject)
	if err != nil {
		return nil, err
	}

	// 判断票据是否为空， 并校验
	if ticket == nil {
		return nil, errors.New("token不存在")
	}

	if !application.tokenService.CheckToken(token, ticket.Secret) {
		return nil, errors.New("令牌校验失败")
	}

	// 校验成功，返回用户id
	return ticket, nil
}

// RefreshToken 刷新令牌
func (application *LoginApplication) RefreshToken(userId int64, subject, token string) (string, error) {
	// 取出票据id
	token, err := application.parseToken(token)
	if err != nil {
		return "", errors.New("token数据不正确")
	}

	// 生成令牌数据
	return application.genericToken(userId)
}

// GetAuthConfigPaths 获取忽略auth的地址，获取普通过滤的地址
func (application *LoginApplication) GetAuthConfigPaths() (ignorePathResultMap map[string]bool,
	ignorePrefixResultPaths []string, commonPathResultMap map[string]bool) {
	if !initPath {
		for _, v := range application.authConfig.IgnorePath {
			ignorePathMap[v] = true
		}

		for _, v := range application.authConfig.IgnorePrefixPath {
			ignorePrefixPaths = append(ignorePrefixPaths, v)
		}

		for _, v := range application.authConfig.CommonPath {
			commonPathMap[v] = true
		}
	}

	initPath = true
	return ignorePathMap, ignorePrefixPaths, commonPathMap
}

// 生成令牌
func (application *LoginApplication) genericToken(userId int64) (string, error) {
	// 生成保存密钥
	secret := uuid.New().String()
	subject := uuid.New().String()

	// 保存用户信息与令牌之间的关系
	// subject -> userId
	// subject -> secret
	// userId -> subject
	err := application.repository.SysTokenRepository.Save(entity.SysToken{
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
	return application.tokenService.GenericToken(application.authConfig, secret, subject)
}

// 从header中解析令牌
func (application *LoginApplication) parseToken(token string) (string, error) {
	// 检查数据
	typeKey := fmt.Sprintf("%s ", application.authConfig.HeaderType)
	typeIndex := strings.Index(token, typeKey)
	if typeIndex != 0 {
		return "", errors.New("token数据不正确")
	}

	// 取出票据id
	return helper.StringHelper.SubString(token, len(typeKey), len(token)), nil
}

// GetUserMenuPermission 用户菜单权限
func (application *LoginApplication) GetUserMenuPermission(userId int64) (*types.SysMenuPermissionForUser, error) {
	sysPermissions, err := application.GetUserSysPermission(userId)
	if err != nil {
		return nil, err
	}

	// 计算菜单id列表, 计算操作id列表
	menuIdMap := make(map[int64]string, 0)
	menuIds := make([]int64, 0)
	opIdMap := make(map[int64]string, 0)
	opIds := make([]int64, 0)
	for _, permission := range sysPermissions {
		_, ok := menuIdMap[permission.MenuId]
		if !ok {
			menuIdMap[permission.MenuId] = ""
			menuIds = append(menuIds, permission.MenuId)
		}

		if permission.Option != "" {
			for _, opStrId := range strings.Split(permission.Option, ",") {
				opId, err := strconv.ParseInt(opStrId, 10, 64)
				if err == nil {
					_, ok := opIdMap[opId]
					if !ok {
						opIdMap[opId] = ""
						opIds = append(opIds, opId)
					}
				}
			}
		}
	}

	// 获取数据组成树
	allMenus, err := application.repository.SysMenuRepository.SelectByIds(menuIds)
	if err != nil {
		return nil, err
	}

	menuCodes := make([]string, 0)
	rootMenus := make([]types.SysMenuPermissionForUserMenu, 0)
	var menuMap = make(map[int64][]types.SysMenuPermissionForUserMenu, 0)
	for _, item := range allMenus {
		// 放code
		menuCodes = append(menuCodes, item.Code)

		// 放数据到menuMap
		menu, ok := menuMap[*item.Parent]
		if !ok {
			menu = make([]types.SysMenuPermissionForUserMenu, 0)
		}

		menu = append(menu, types.SysMenuPermissionForUserMenu{
			Id:        item.Id,
			Icon:      item.Icon,
			Component: item.Component,
			Path:      item.Path,
			Name:      item.Code,
		})
		menuMap[*item.Parent] = menu

		// 得到rootMenus
		if *item.Parent == 0 {
			rootMenus = append(rootMenus, types.SysMenuPermissionForUserMenu{
				Id:        item.Id,
				Icon:      item.Icon,
				Component: item.Component,
				Path:      item.Path,
				Name:      item.Code,
			})
		}
	}

	application.recursionAssignmentForUserMenu(rootMenus, menuMap)

	// 获取操作组成树
	menuOptions, err := application.repository.SysMenuOptionRepository.SelectByIds(opIds)
	if err != nil {
		return nil, err
	}

	optionValues := make([]string, 0)
	for _, option := range menuOptions {
		optionValues = append(optionValues, option.Value)
	}
	return &types.SysMenuPermissionForUser{
		Permissions: optionValues,
		Menus:       rootMenus,
		Codes:       menuCodes,
	}, nil
}

// GetUserMenuUrl 获取用户拥有的权限路径
func (application *LoginApplication) GetUserMenuUrl(userId int64) (map[string]bool, error) {
	specMap := make(map[string]bool, 0)
	options, err := application.GetUserSysMenuOption(userId)
	if err != nil {
		return specMap, err
	}

	for _, option := range options {
		fullKey := fmt.Sprintf("%s - %s", option.Method, option.Path)
		specMap[fullKey] = true
	}
	return specMap, err
}

// GetUserSysMenuOption 获取用户拥有的路径
func (application *LoginApplication) GetUserSysMenuOption(userId int64) ([]entity.SysMenuOption, error) {
	sysPermissions, err := application.GetUserSysPermission(userId)
	if err != nil {
		return nil, err
	}

	// 计算菜单id列表, 计算操作id列表
	opIdMap := make(map[int64]string, 0)
	opIds := make([]int64, 0)
	for _, permission := range sysPermissions {
		if permission.Option != "" {
			for _, opStrId := range strings.Split(permission.Option, ",") {
				opId, err := strconv.ParseInt(opStrId, 10, 64)
				if err == nil {
					_, ok := opIdMap[opId]
					if !ok {
						opIdMap[opId] = ""
						opIds = append(opIds, opId)
					}
				}
			}
		}
	}
	// 获取操作组成树
	return application.repository.SysMenuOptionRepository.SelectByIds(opIds)
}

// GetUserSysPermission 获取用户所拥有的权限列表
func (application *LoginApplication) GetUserSysPermission(userId int64) ([]entity.SysPermission, error) {
	sysUser, err := application.repository.SysUserRepository.GetById(userId)
	if err != nil {
		return nil, err
	}

	if sysUser.Roles == "" {
		return nil, errors.New("角色信息不存在")
	}

	// 获取角色id列表
	roleIds := make([]int64, 0)
	roleStrIds := strings.Split(sysUser.Roles, ",")
	for _, roleStrId := range roleStrIds {
		roleId, err := strconv.ParseInt(roleStrId, 10, 64)
		if err != nil {
			return nil, errors.New("角色信息有误")
		}
		roleIds = append(roleIds, roleId)
	}

	// 角色id列表到permission表中查询
	return application.repository.SysPermissionRepository.SelectByRoleIds(roleIds)
}

func (application *LoginApplication) recursionAssignmentForUserMenu(
	rootMenus []types.SysMenuPermissionForUserMenu,
	menuMap map[int64][]types.SysMenuPermissionForUserMenu) {

	for index, item := range rootMenus {
		menus, _ := menuMap[item.Id]
		item.Routes = menus
		rootMenus[index] = item

		// 判断退出条件
		if menus != nil && len(menus) != 0 {
			// 继续赋值
			application.recursionAssignmentForUserMenu(menus, menuMap)
		}
	}
}
