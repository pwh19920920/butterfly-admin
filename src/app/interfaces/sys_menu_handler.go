package interfaces

import (
	"butterfly-admin/src/app/application"
	"butterfly-admin/src/app/types"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
)

type sysMenuHandler struct {
	menuApp application.SysMenuApplication
}

// 查询
func (handler *sysMenuHandler) query(context *gin.Context) {
	var sysMenuQueryRequest types.SysMenuQueryRequest
	if context.ShouldBindQuery(&sysMenuQueryRequest) != nil {
		response.BuildResponseBadRequest(context, "请求参数有误")
		return
	}

	// option
	total, data, err := handler.menuApp.Query(sysMenuQueryRequest)
	if err != nil {
		response.BuildResponseBadRequest(context, "用户名或者密码错误")
		return
	}

	// 输出
	response.BuildResponseSuccess(context, response.RespPaging{
		RespBody: response.GenericResponse(200, "", data),
		PageSize: sysMenuQueryRequest.GetPageSize(),
		Total:    total,
		Current:  sysMenuQueryRequest.GetCurrent(),
	})
}

// 加载路由
func InitSysMenuHandler(app *application.Application) {
	// 组件初始化
	handler := sysMenuHandler{app.SysMenu}

	// 路由初始化
	var route []server.RouteInfo
	route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "", HandlerFunc: handler.query})
	server.RegisterRoute("/sys/menu", route)
}
