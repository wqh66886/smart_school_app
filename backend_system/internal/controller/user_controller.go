package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wqh/smart/school/system/internal/initiate"
	"github.com/wqh/smart/school/system/internal/service"
	"github.com/wqh/smart/school/system/response"
	"math/rand"
	"time"
)

/**
* description:
* author: wqh
* date: 2025/1/8
 */
type UserController struct {
	userSrv   *service.UserService
	schoolSrv *service.SchoolService
}

func NewUserController(userSrv *service.UserService, schoolSrv *service.SchoolService) *UserController {
	return &UserController{userSrv: userSrv, schoolSrv: schoolSrv}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var registerInfo response.RegisterInfo
	if err := ctx.ShouldBindJSON(&registerInfo); err != nil {
		ctx.Error(initiate.INVALID_ARGUMENT)
		return
	}
	if err := uc.userSrv.Register(registerInfo); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, response.Response{
		Code:    200,
		Message: "注册成功",
		Data:    nil,
	})
}

func (uc *UserController) SendVerificationCode(ctx *gin.Context) {
	phone := ctx.Query("phone")
	if len(phone) == 0 {
		ctx.Error(initiate.INNER_ERROR)
		return
	}
	//:TODO 短信端口接入
	code := rand.Intn(900000) + 100000
	initiate.RDB.Set(ctx, phone, code, time.Minute*5)
	ctx.JSON(200, response.Response{
		Code:    200,
		Message: "发送成功",
		Data:    nil,
	})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var loginInfo response.LoginInfo
	if err := ctx.ShouldBindJSON(&loginInfo); err != nil {
		ctx.Error(initiate.INVALID_ARGUMENT)
		return
	}
	token, err := uc.userSrv.Login(loginInfo)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, response.Response{
		Code:    200,
		Message: "登录成功",
		Data: gin.H{
			"token": token,
		},
	})
}
