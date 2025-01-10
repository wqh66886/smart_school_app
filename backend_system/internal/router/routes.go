package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh/smart/school/system/internal/router/user"
)

func InitRouter(engine *gin.Engine) {
	public := engine.Group("v1")
	user.UserRouter(public)
}
