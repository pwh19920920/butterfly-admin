package interfaces

import (
	"butterfly-admin/src/app/application"
	"butterfly-admin/src/app/common/constant"
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/types"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
)

type sysUserHandler struct {
	userApp application.SysUserApplication
}

// 登陆
func (userHandler *sysUserHandler) login(context *gin.Context) {
	var form types.LoginForm
	if context.ShouldBindJSON(&form) != nil {
		response.BuildResponseBadRequest(context, "请求参数有误")
		return
	}

	// option
	token, err := userHandler.userApp.Login(form.Username, form.Password)
	if err != nil {
		response.BuildResponseBadRequest(context, "用户名或者密码错误")
		return
	}

	// 输出
	response.BuildResponseSuccess(context, token)
}

// 退出
func (userHandler *sysUserHandler) logout(context *gin.Context) {
	// 尝试获取ticket
	dataStr := context.Request.Header.Get(constant.ContextUser)
	token := entity.SysToken{}.UnMarshal(dataStr)

	// 删除令牌
	_ = userHandler.userApp.Logout(token.RelationId)

	// 输出
	response.BuildResponseSuccess(context, token)
}

// 刷新令牌
func (userHandler *sysUserHandler) refresh(context *gin.Context) {
	// 尝试获取ticket
	dataStr := context.Request.Header.Get(constant.ContextUser)
	ticket := entity.SysToken{}.UnMarshal(dataStr)

	// 取令牌
	token := context.GetHeader(userHandler.userApp.GetHeaderName())
	newToken, err := userHandler.userApp.RefreshToken(ticket.UserId, ticket.RelationId, token)
	if err != nil {
		response.BuildResponseBadRequest(context, "刷新令牌失败")
		return
	}

	// 删除令牌
	_ = userHandler.userApp.Logout(ticket.RelationId)

	// 输出
	response.BuildResponseSuccess(context, newToken)
}

// 加载路由
func InitSysUserHandler(app *application.Application) {
	// 组件初始化
	handler := sysUserHandler{app.SysUser}

	// 路由初始化
	var route []server.RouteInfo
	route = append(route, server.RouteInfo{HttpMethod: server.HttpPost, Path: "/login", HandlerFunc: handler.login})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpPost, Path: "/logout", HandlerFunc: handler.logout})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpPost, Path: "/refresh", HandlerFunc: handler.refresh})
	server.RegisterRoute("/sys", route)
}
