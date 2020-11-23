package middleware

import (
	"butterfly-admin/src/app/application"
	"butterfly-admin/src/app/common/constant"
	"fmt"
	"github.com/gin-gonic/gin"
)

func JwtAuth(app *application.Application, routeFor401 gin.HandlerFunc, routeFor403 gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 404了，不处理
		if context.FullPath() == "" {
			return
		}

		// 如果地址被忽略
		urlFullKey := fmt.Sprintf("%s - %s", context.Request.Method, context.FullPath())
		ignorePaths := app.SysUser.GetIgnorePaths()
		if ignorePaths != nil {
			_, ok := (*ignorePaths)[urlFullKey]
			if ok {
				context.Next()
				return
			}
		}

		// 不被忽略则判断是否有令牌
		token := context.GetHeader(app.SysUser.GetHeaderName())
		ticket, err := app.SysUser.CheckAndGetTicket(token)
		if err != nil {
			routeFor401(context)
			context.Abort()
			return
		}

		// 特殊权限校验
		specUserPermission := app.SysUser.GetUserPermission(ticket.UserId)
		_, ok := (*specUserPermission)[urlFullKey]
		if !ok {
			routeFor403(context)
			context.Abort()
			return
		}

		// 设置用户的基本信息
		context.Request.Header.Set(constant.ContextUser, ticket.Marshal())

		// 如果需要权限则继续判断
		context.Next()
	}
}
