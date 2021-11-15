package starter

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly-admin/application"
	"github.com/pwh19920920/butterfly-admin/config"
	"github.com/pwh19920920/butterfly-admin/infrastructure/persistence"
	"github.com/pwh19920920/butterfly-admin/infrastructure/security"
	"github.com/pwh19920920/butterfly-admin/interfaces"
	"github.com/pwh19920920/butterfly-admin/interfaces/middleware"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
)

func route401(context *gin.Context) {
	response.Response(context, 401, "请登录后在进行此操作", nil)
}

func route403(context *gin.Context) {
	response.Response(context, 403, "您没有权限进行此操作", nil)
}

func InitButterflyAdmin() config.Config {
	// 初始化基本服务
	allConfig := config.InitAll()
	repository := persistence.NewRepository(allConfig)
	encodeService := security.NewEncodeServiceImpl()
	tokenService := security.NewJwtServiceImpl()
	app := application.NewApplication(
		allConfig,
		repository,
		encodeService,
		tokenService,
	)

	// 初始化相关路由
	interfaces.InitLoginHandler(app)
	interfaces.InitSysMenuHandler(app)
	interfaces.InitSysRoleHandler(app)
	interfaces.InitSysUserHandler(app)

	// 注册中间对象
	server.RegisterMiddleware(middleware.JwtAuth(
		app,
		route401,
		route403,
	))
	return allConfig
}
