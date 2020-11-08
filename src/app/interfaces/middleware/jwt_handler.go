package middleware

import (
	"butterfly-admin/src/app/application"
	"butterfly-admin/src/app/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

func JwtAuth(user *application.User, routeFor401 gin.HandlerFunc, routeFor403 gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 404了，不处理
		if context.FullPath() == "" {
			return
		}

		// 如果地址被忽略
		ignorePaths := user.GetIgnorePaths()
		if ignorePaths != nil {
			_, ok := (*ignorePaths)[fmt.Sprintf("%s - %s", context.Request.Method, context.FullPath())]
			if ok {
				context.Next()
				return
			}
		}

		// 不被忽略则判断是否有令牌
		token := context.GetHeader(user.GetHeaderName())
		userId, err := user.CheckAndGetUserId(token)
		if err != nil {
			routeFor401(context)
			context.Abort()
			return
		}

		// 设置用户的基本信息
		context.Header(common.ContextUser, userId)

		// 如果需要权限则继续判断
		context.Next()
	}
}
