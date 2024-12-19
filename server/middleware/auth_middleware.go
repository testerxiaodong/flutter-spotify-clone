package middleware

import (
	"server/common/response"
	"server/common/xerror"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("x-auth-token")
		if len(tokenString) == 0 {
			context.Abort()
			response.Result(context, nil, xerror.NewErrCodeMsg(xerror.AUTH_CHECK_FAILURE, "用户未登录"))
			return
		}
		userId, err := utils.ParseToken(tokenString)
		if err != nil {
			context.Abort()
			response.Result(context, nil, xerror.NewErrCodeMsg(xerror.AUTH_CHECK_FAILURE, "token无效"))
			return
		}
		context.Set("userId", userId)
		context.Next()
	}
}
