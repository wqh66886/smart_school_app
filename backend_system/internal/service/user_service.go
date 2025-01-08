package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/wqh/smart/school/system/internal/domain"
	"github.com/wqh/smart/school/system/internal/errorx"
	"github.com/wqh/smart/school/system/internal/initiate"
	"github.com/wqh/smart/school/system/internal/repository"
	"github.com/wqh/smart/school/system/internal/utils"
	"github.com/wqh/smart/school/system/response"
	"time"
)

/**
* description:
* author: wqh
* date: 2025/1/8
 */
type UserService struct {
	UserRepository   repository.UserRepository
	SchoolRepository repository.SchoolRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (u *UserService) Register(userInfo response.RegisterInfo) error {
	if len(userInfo.Phone) == 0 || len(userInfo.Password) == 0 {
		return errorx.NewError(400, "手机号或密码不能为空")
	}
	user, err := u.UserRepository.GetUserByPhone(userInfo.Phone)
	if err != nil {
		return initiate.INNER_ERROR
	}
	if user != nil {
		return errorx.NewError(400, "手机号已存在")
	}
	pwd, err := utils.GetPwd(userInfo.Password)
	if err != nil {
		return initiate.INNER_ERROR
	}
	user = &domain.User{
		Base: domain.Base{
			Id:         uuid.NewV4().String(),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
		Phone:    userInfo.Phone,
		Password: pwd,
	}
	if err := u.UserRepository.CreateUser(user); err != nil {
		return errorx.NewError(500, "内部错误")
	}
	return nil
}

func (u *UserService) Login(info response.LoginInfo) (string, error) {
	return "", nil
}

func (u *UserService) UpdateUserInfo(userInfo response.RegisterInfo) error {
	if len(userInfo.Phone) == 0 {
		return initiate.INNER_ERROR
	}
	if len(userInfo.SchoolCode) == 0 {
		return errorx.NewError(400, "学习必须选择!")
	}
	user, err := u.UserRepository.GetUserByPhone(userInfo.Phone)
	if err != nil {
		return initiate.INNER_ERROR
	}
	if user == nil {
		return errorx.NewError(400, "用户不存在")
	}
	school, err := u.SchoolRepository.GetSchoolByCode(userInfo.SchoolCode)
	if err != nil {
		return initiate.INNER_ERROR
	}
	user.SchoolId = school.Id
	user.Username = userInfo.Username
	user.UpdateTime = time.Now()
	user.Avatar = userInfo.Avatar
	user.Gender = userInfo.Gender
	user.Birthday = utils.ParseTime(userInfo.Birthday)
	if err := u.UserRepository.UpdateUserInfo(user); err != nil {
		return initiate.INNER_ERROR
	}
	return nil
}
