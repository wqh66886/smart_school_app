package middleware

import (
	"github.com/wqh/smart/school/system/internal/errorx"
	"github.com/wqh/smart/school/system/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()               // 调用后续中间件
		if len(ctx.Errors) > 0 { // 判断是否有错误
			e := ctx.Errors.Last() // 获取最后一个错误
			if err, ok := e.Err.(*errorx.Error); ok {
				ctx.JSON(err.Code, response.Response{
					Code:    err.Code,
					Message: err.Message,
					Data:    err.Data,
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, response.Response{
					Code:    http.StatusInternalServerError,
					Message: e.Err.Error(),
					Data:    nil,
				})
			}
			ctx.Abort()
			return
		}
	}
}
