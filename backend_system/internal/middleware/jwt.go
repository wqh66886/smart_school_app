package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh/smart/school/system/internal/initiate"
	"github.com/wqh/smart/school/system/internal/utils"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if len(token) == 0 {
			ctx.Error(initiate.UNAUTHENTICATED)
			ctx.Abort()
			return
		}
		claims, err := utils.IsAuthorized(token)
		if err != nil || claims == nil {
			ctx.Error(initiate.UNAUTHENTICATED)
			ctx.Abort()
			return
		}

		ctx.Set("userId", claims.ID)
		ctx.Set("userName", claims.Name)
		ctx.Next()
	}
}
