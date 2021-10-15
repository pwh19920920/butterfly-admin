package application

import (
	"github.com/pwh19920920/butterfly-admin/src/app/config"
	"github.com/pwh19920920/butterfly-admin/src/app/domain/security"
	"github.com/pwh19920920/butterfly-admin/src/app/infrastructure/persistence"
)

type Application struct {
	Login         LoginApplication
	SysMenu       SysMenuApplication
	SysUser       SysUserApplication
	SysRole       SysRoleApplication
	SysPermission SysPermissionApplication
}

func NewApplication(
	config config.Config,
	repository *persistence.Repository,
	encoderService security.EncodeService,
	tokenService security.TokenService,
) *Application {
	return &Application{
		// 用户服务
		Login: LoginApplication{
			sequence:       config.Sequence,
			repository:     repository,
			encoderService: encoderService,
			tokenService:   tokenService,
			authConfig:     config.AuthConfig,
		},

		// 菜单服务
		SysMenu: SysMenuApplication{
			sequence:   config.Sequence,
			repository: repository,
		},

		SysUser: SysUserApplication{
			sequence:       config.Sequence,
			repository:     repository,
			encoderService: encoderService,
			tokenService:   tokenService,
			authConfig:     config.AuthConfig,
		},

		// 角色
		SysRole: SysRoleApplication{
			sequence:   config.Sequence,
			repository: repository,
		},

		// 权限
		SysPermission: SysPermissionApplication{
			sequence:   config.Sequence,
			repository: repository,
		},
	}
}
