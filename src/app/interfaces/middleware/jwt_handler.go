package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly-admin/src/app/application"
	"github.com/pwh19920920/butterfly-admin/src/app/common/constant"
	"strings"
)

func JwtAuth(app *application.Application, routeFor401 gin.HandlerFunc, routeFor403 gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 404了，不处理
		if context.FullPath() == "" {
			return
		}

		// 如果地址被忽略
		urlFullKey := fmt.Sprintf("%s - %s", context.Request.Method, context.FullPath())

		// 先判断前缀过滤
		ignorePaths, ignorePrefixPaths, commonPathMap := app.Login.GetAuthConfigPaths()
		if ignorePrefixPaths != nil && len(ignorePrefixPaths) > 0 {
			for _, path := range ignorePrefixPaths {
				if strings.HasPrefix(urlFullKey, path) {
					context.Next()
					return
				}
			}
		}

		// 判断全匹配过滤
		_, ok := ignorePaths[urlFullKey]
		if ok {
			context.Next()
			return
		}

		// 不被忽略则判断是否有令牌
		token := context.GetHeader(app.Login.GetHeaderName())
		ticket, err := app.Login.CheckAndGetTicket(token)
		if err != nil {
			routeFor401(context)
			context.Abort()
			return
		}

		// 判断是否有普通权限配置匹配
		_, ok = commonPathMap[urlFullKey]
		if ok {
			// 设置用户的基本信息
			context.Request.Header.Set(constant.ContextUser, ticket.Marshal())
			context.Next()
			return
		}

		// 特殊权限校验
		specUserPermission, err := app.Login.GetUserMenuUrl(ticket.UserId)
		if err != nil {
			routeFor403(context)
			context.Abort()
			return
		}

		// 判断是否存在, 存在说明有权限
		_, ok = specUserPermission[urlFullKey]
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
