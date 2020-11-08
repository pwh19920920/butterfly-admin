package main

import (
	"butterfly-admin/src/app/application"
	"butterfly-admin/src/app/config/auth"
	"butterfly-admin/src/app/config/database"
	"butterfly-admin/src/app/infrastructure/persistence"
	"butterfly-admin/src/app/interfaces"
	"butterfly-admin/src/app/interfaces/middleware"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
)

func route401(context *gin.Context) {
	response.Response(context, 401, "请登录后在进行此操作", nil)
}

func route403(context *gin.Context) {
	response.Response(context, 403, "您没有权限进行此操作", nil)
}

func init() {
	// 初始化持久层
	db := database.GetConn()
	repository := persistence.NewRepository(db)

	// 初始化相关路由
	interfaces.InitUserHandler(repository, auth.GetConf())

	// 注册中间对象
	server.RegisterMiddleware(middleware.JwtAuth(
		application.NewUser(repository, auth.GetConf()),
		route401,
		route403,
	))
}

func main() {
	butterfly.Run()
}
