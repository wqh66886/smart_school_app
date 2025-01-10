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

// Register 用户注册
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

// SendVerificationCode 发送验证码
func (uc *UserController) SendVerificationCode(ctx *gin.Context) {
	smsType := ctx.Query("smsType")
	if len(smsType) == 0 {
		ctx.Error(initiate.INNER_ERROR)
		return
	}
	code := rand.Intn(900000) + 100000
	var key string
	if smsType == "phone" {
		//:TODO 短信端口接入
		key = ctx.Query("phone")
	} else {
		key = ctx.Query("email")
	}
	initiate.RDB.Set(ctx, key, code, time.Minute*5)
	ctx.JSON(200, response.Response{
		Code:    200,
		Message: "发送成功",
		Data:    nil,
	})
}

// UpdateUserInfo 修改用户信息
func (uc *UserController) UpdateUserInfo(ctx *gin.Context) {
	var userInfo response.RegisterInfo
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.Error(initiate.INVALID_ARGUMENT)
		return
	}
	if len(userInfo.SchoolCode) == 0 {
		ctx.Error(initiate.INVALID_ARGUMENT)
		return
	}
	if school, err := uc.schoolSrv.GetSchoolInfoByCode(userInfo.SchoolCode); err != nil {
		ctx.Error(err)
		return
	} else if err = uc.userSrv.UpdateUserInfo(userInfo, school); err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, response.Response{
		Code: 200,
	})
}

// Login 用户登录
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
