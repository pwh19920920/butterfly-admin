package interfaces

import (
	"butterfly-admin/src/app/application"
	"butterfly-admin/src/app/common/constant"
	"butterfly-admin/src/app/config"
	"butterfly-admin/src/app/domain/entity"
	"butterfly-admin/src/app/infrastructure/persistence"
	"butterfly-admin/src/app/types"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
)

type userHandler struct {
	oauth *application.User
}

func (userHandler *userHandler) login(context *gin.Context) {
	var form types.LoginForm
	if context.ShouldBindJSON(&form) != nil {
		response.BuildResponseBadRequest(context, "请求参数有误")
		return
	}

	// option
	token, err := userHandler.oauth.Login(form.Username, form.Password)
	if err != nil {
		response.BuildResponseBadRequest(context, "用户名或者密码错误")
		return
	}

	// 输出
	response.BuildResponseSuccess(context, token)
}

func (userHandler *userHandler) logout(context *gin.Context) {
	// 尝试获取ticket
	dataStr := context.Request.Header.Get(constant.ContextUser)
	token := entity.Token{}.UnMarshal(dataStr)

	// 删除令牌
	_ = userHandler.oauth.Logout(token.RelationId)

	// 输出
	response.BuildResponseSuccess(context, token)
}

// 加载路由
func InitUserHandler(repository *persistence.Repository, authConfig *config.AuthConfig) {
	// 组件初始化
	oauth := application.NewUser(repository, authConfig)
	handler := userHandler{oauth}

	// 路由初始化
	var route []server.RouteInfo
	route = append(route, server.RouteInfo{HttpMethod: server.HttpPost, Path: "/login", HandlerFunc: handler.login})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpPost, Path: "/logout", HandlerFunc: handler.logout})
	server.RegisterRoute("/sys", route)
}
