package interfaces

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly-admin/application"
	"github.com/pwh19920920/butterfly-admin/types"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
	"strconv"
)

type sysMenuHandler struct {
	menuApp application.SysMenuApplication
}

// 查询
func (handler *sysMenuHandler) queryWithoutOption(context *gin.Context) {
	// option
	data, err := handler.menuApp.QueryForTree(false)
	if err != nil {
		response.BuildResponseBadRequest(context, "查询菜单出错")
		return
	}

	// 输出
	response.BuildResponseSuccess(context, data)
}

// 查询带op
func (handler *sysMenuHandler) queryWithOption(context *gin.Context) {
	// option
	data, err := handler.menuApp.QueryForTree(true)
	if err != nil {
		response.BuildResponseBadRequest(context, "查询菜单出错")
		return
	}

	// 输出
	response.BuildResponseSuccess(context, data)
}

// 创建
func (handler *sysMenuHandler) create(context *gin.Context) {
	var sysMenuCreateRequest types.SysMenuCreateRequest
	if context.ShouldBindJSON(&sysMenuCreateRequest) != nil {
		response.BuildResponseBadRequest(context, "请求参数有误")
		return
	}

	// 数据校验
	err := sysMenuCreateRequest.ValidateForCreate()
	if err != nil {
		response.BuildResponseBadRequest(context, fmt.Sprintf("请求参数有误: %v", err.Error()))
		return
	}

	// option
	err = handler.menuApp.Create(&sysMenuCreateRequest)
	if err != nil {
		response.BuildResponseBadRequest(context, "创建菜单失败")
		return
	}

	response.BuildResponseSuccess(context, "ok")
}

// 更新
func (handler *sysMenuHandler) update(context *gin.Context) {
	var sysMenuCreateRequest types.SysMenuCreateRequest
	err := context.ShouldBindJSON(&sysMenuCreateRequest)
	if err != nil {
		response.BuildResponseBadRequest(context, "请求参数有误")
		return
	}

	// 数据校验
	err = sysMenuCreateRequest.ValidateForModify()
	if err != nil {
		response.BuildResponseBadRequest(context, fmt.Sprintf("请求参数有误: %v", err.Error()))
		return
	}

	// option
	err = handler.menuApp.Modify(&sysMenuCreateRequest)
	if err != nil {
		response.BuildResponseBadRequest(context, "更新菜单失败")
		return
	}

	response.BuildResponseSuccess(context, "ok")
}

// 删除
func (handler *sysMenuHandler) delete(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BuildResponseBadRequest(context, "请求参数有无")
		return
	}

	// option
	err = handler.menuApp.Delete(id)
	if err != nil {
		response.BuildResponseBadRequest(context, "删除菜单失败:"+err.Error())
		return
	}

	response.BuildResponseSuccess(context, "ok")
}

// 删除
func (handler *sysMenuHandler) option(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BuildResponseBadRequest(context, "请求参数有无")
		return
	}

	// option
	data, err := handler.menuApp.QueryOptionByMenuId(id)
	if err != nil {
		response.BuildResponseBadRequest(context, "获取操作失败")
		return
	}

	response.BuildResponseSuccess(context, data)
}

// InitSysMenuHandler 加载路由
func InitSysMenuHandler(app *application.Application) {
	// 组件初始化
	handler := sysMenuHandler{app.SysMenu}

	// 路由初始化
	var route []server.RouteInfo
	route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "", HandlerFunc: handler.queryWithoutOption})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "/withOption", HandlerFunc: handler.queryWithOption})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "/option/:id", HandlerFunc: handler.option})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpPost, Path: "", HandlerFunc: handler.create})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpPut, Path: "", HandlerFunc: handler.update})
	route = append(route, server.RouteInfo{HttpMethod: server.HttpDelete, Path: "/:id", HandlerFunc: handler.delete})
	server.RegisterRoute("/api/sys/menu", route)
}
