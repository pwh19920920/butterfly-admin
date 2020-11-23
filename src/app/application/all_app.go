package application

import (
	"butterfly-admin/src/app/config"
	"butterfly-admin/src/app/domain/security"
	"butterfly-admin/src/app/infrastructure/persistence"
	"github.com/bwmarrin/snowflake"
)

type Application struct {
	SysUser SysUserApplication
}

func NewApplication(
	sequence *snowflake.Node,
	repository *persistence.Repository,
	encoderService security.EncodeService,
	jwtService security.JwtService,
	authConfig *config.AuthConfig,
) *Application {
	return &Application{
		SysUser: SysUserApplication{
			sequence:       sequence,
			repository:     repository,
			encoderService: encoderService,
			jwtService:     jwtService,
			authConfig:     authConfig,
		},
	}
}
