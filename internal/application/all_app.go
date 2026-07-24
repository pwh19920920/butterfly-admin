package application

import (
	"github.com/pwh19920920/butterfly-admin/internal/config"
	"github.com/pwh19920920/butterfly-admin/internal/config/auth"
	"github.com/pwh19920920/butterfly-admin/internal/domain/security"
	"github.com/pwh19920920/butterfly-admin/internal/infrastructure/persistence"
	"github.com/pwh19920920/snowflake"
)

// baseApp 应用层公共依赖，通过结构体嵌入复用到各个 Application
type baseApp struct {
	sequence       *snowflake.Node
	repository     *persistence.Repository
	encoderService security.EncodeService
	tokenService   security.TokenService
	authConfig     *auth.Config
}

type Application struct {
	Login         LoginApplication
	SysMenu       SysMenuApplication
	SysUser       SysUserApplication
	SysRole       SysRoleApplication
	SysPermission SysPermissionApplication
}

func NewApplication(
	cfg config.Config,
	repository *persistence.Repository,
	encoderService security.EncodeService,
	tokenService security.TokenService,
) *Application {
	base := baseApp{
		sequence:       cfg.Sequence,
		repository:     repository,
		encoderService: encoderService,
		tokenService:   tokenService,
		authConfig:     cfg.AuthConfig,
	}

	return &Application{
		Login:         LoginApplication{baseApp: base},
		SysMenu:       SysMenuApplication{baseApp: base},
		SysUser:       SysUserApplication{baseApp: base},
		SysRole:       SysRoleApplication{baseApp: base},
		SysPermission: SysPermissionApplication{baseApp: base},
	}
}
