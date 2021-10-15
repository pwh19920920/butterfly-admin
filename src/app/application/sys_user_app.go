package application

import (
	"butterfly-admin/src/app/config/auth"
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/domain/security"
	"butterfly-admin/src/app/infrastructure/persistence"
	"butterfly-admin/src/app/types"
	"github.com/bwmarrin/snowflake"
	"github.com/go-basic/uuid"
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

func (app *SysUserApplication) Query(request *types.SysUserQueryRequest) (int64, []types.SysUserQueryResponse, error) {
	total, data, err := app.repository.SysUserRepository.Select(request)
	// 错误记录
	if err != nil {
		logrus.Error("SysUserRepository.Select() happen error for", err)
		return total, nil, err
	}

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
		})
	}
	return total, result, err
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
