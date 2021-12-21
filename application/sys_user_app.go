package application

import (
	"github.com/go-basic/uuid"
	"github.com/pwh19920920/butterfly-admin/config/auth"
	"github.com/pwh19920920/butterfly-admin/domain/entity"
	"github.com/pwh19920920/butterfly-admin/domain/security"
	"github.com/pwh19920920/butterfly-admin/infrastructure/persistence"
	"github.com/pwh19920920/butterfly-admin/types"
	"github.com/pwh19920920/snowflake"
	"github.com/sirupsen/logrus"
	"strings"
)

type SysUserApplication struct {
	sequence       *snowflake.Node
	repository     *persistence.Repository
	encoderService security.EncodeService
	tokenService   security.TokenService
	authConfig     *auth.Config
}

func (app *SysUserApplication) GetUserById(userId int64) (*entity.SysUser, error) {
	return app.repository.SysUserRepository.GetById(userId)
}

func (app *SysUserApplication) coverQueryResult(data []entity.SysUser) []types.SysUserQueryResponse {
	// 重新赋值
	result := make([]types.SysUserQueryResponse, 0)
	for _, item := range data {
		roleList := make([]string, 0)
		if item.Roles != "" {
			roleList = strings.Split(item.Roles, ",")
		}
		result = append(result, types.SysUserQueryResponse{
			BaseEntity: item.BaseEntity,
			Name:       item.Name,
			Avatar:     item.Avatar,
			Roles:      item.Roles,
			Username:   item.Username,
			RoleList:   roleList,
			Email:      item.Email,
			Mobile:     item.Mobile,
		})
	}
	return result
}

func (app *SysUserApplication) Query(request *types.SysUserQueryRequest) (int64, []types.SysUserQueryResponse, error) {
	total, data, err := app.repository.SysUserRepository.Select(request)
	// 错误记录
	if err != nil {
		logrus.Error("SysUserRepository.Select() happen error for", err)
		return total, nil, err
	}

	// 重新赋值
	result := app.coverQueryResult(data)
	return total, result, err
}

// QueryAll 查询全部
func (app *SysUserApplication) QueryAll() ([]types.SysUserQueryResponse, error) {
	data, err := app.repository.SysUserRepository.SelectAll()
	// 重新赋值
	result := app.coverQueryResult(data)
	return result, err
}

// Create 创建用户
func (app *SysUserApplication) Create(user *entity.SysUser) error {
	user.Id = app.sequence.Generate().Int64()
	user.Salt = uuid.New()
	user.Password = app.encoderService.Encode(user.Password, user.Salt)
	return app.repository.SysUserRepository.Create(user)
}

// Modify 更新用户
func (app *SysUserApplication) Modify(user *entity.SysUser) error {
	if user.Password != "" {
		user.Salt = uuid.New()
		user.Password = app.encoderService.Encode(user.Password, user.Salt)
	}
	return app.repository.SysUserRepository.Modify(user)
}
