package user

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh/smart/school/system/internal/controller"
	"github.com/wqh/smart/school/system/internal/initiate"

	"github.com/wqh/smart/school/system/internal/service"
	"github.com/wqh/smart/school/system/internal/usecase"
)

func UserRouter(group *gin.RouterGroup) {
	usersrv := service.NewUserService(usecase.NewUserUseCase(initiate.DB))
	schoolsrv := service.NewSchoolService(usecase.NewSchoolUseCase(initiate.DB))
	uc := controller.NewUserController(usersrv, schoolsrv)
	group.POST("/register", uc.Register)
	group.POST("/login", uc.Login)
	group.POST("/send_verification_code", uc.SendVerificationCode)
	group.POST("/update_user_info", uc.UpdateUserInfo)
}
