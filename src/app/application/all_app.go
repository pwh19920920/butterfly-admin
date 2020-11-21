package application

import (
	"butterfly-admin/src/app/config"
	"butterfly-admin/src/app/domain/security"
	"butterfly-admin/src/app/infrastructure/persistence"
)

type Application struct {
	User UserApplication
}

func NewApplication(
	repository *persistence.Repository,
	encoderService security.EncodeService,
	jwtService security.JwtService,
	authConfig *config.AuthConfig,
) *Application {
	return &Application{
		User: UserApplication{
			repository:     repository,
			encoderService: encoderService,
			jwtService:     jwtService,
			authConfig:     authConfig,
		},
	}
}
