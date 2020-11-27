package application

import (
	"butterfly-admin/src/app/config"
	"butterfly-admin/src/app/domain/security"
	"butterfly-admin/src/app/infrastructure/persistence"
	"github.com/bwmarrin/snowflake"
)

type Application struct {
	SysUser SysUserApplication
	SysMenu SysMenuApplication
}

func NewApplication(
	sequence *snowflake.Node,
	repository *persistence.Repository,
	encoderService security.EncodeService,
	jwtService security.JwtService,
	authConfig *config.AuthConfig,
) *Application {
	return &Application{
		// 用户服务
		SysUser: SysUserApplication{
			sequence:       sequence,
			repository:     repository,
			encoderService: encoderService,
			jwtService:     jwtService,
			authConfig:     authConfig,
		},

		// 菜单服务
		SysMenu: SysMenuApplication{
			sequence:   sequence,
			repository: repository,
		},
	}
}
