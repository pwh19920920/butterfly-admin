package interfaces

import (
	"butterfly-admin/src/app/application"
	"butterfly-admin/src/app/types"
	"fmt"
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
	total, data, err := handler.menuApp.Query(&sysMenuQueryRequest)
	if err != nil {
		response.BuildResponseBadRequest(context, "用户名或者密码错误")
		return
	}

	// 输出
	response.BuildPageResponseSuccess(context, sysMenuQueryRequest.RequestPaging, total, data)
}

// 创建
func (handler *sysMenuHandler) create(context *gin.Context) {
	var sysMenuCreateRequest types.SysMenuCreateRequest
	if context.ShouldBindJSON(&sysMenuCreateRequest) != nil {
		response.BuildResponseBadRequest(context, "请求参数有误")
		return
	}

	// option
	err := handler.menuApp.Create(&sysMenuCreateRequest)
	if err != nil {
		fmt.Printf("%v", err)
		response.BuildResponseBadRequest(context, "创建菜单失败")
		return
	}

	response.BuildResponseSuccess(context, "ok")
}

// 加载路由
func InitSysMenuHandler(app *application.Application) {
	// 组件初始化
	handler := sysMenuHandler{app.SysMenu}

	// 路由初始化
	var route []server.RouteInfo
	route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "", HandlerFunc: handler.query})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpPost, Path: "", HandlerFunc: handler.create})
	server.RegisterRoute("/sys/menu", route)
}
