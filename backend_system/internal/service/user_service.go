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
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Register(userInfo response.RegisterInfo) error {
	if len(userInfo.RegisterType) == 0 {
		return initiate.INVALID_ARGUMENT
	}

	if userInfo.RegisterType == "email" {
		if len(userInfo.Email) == 0 || len(userInfo.Password) == 0 {
			return errorx.NewError(400, "邮箱或密码不能为空")
		}
		user, err := u.userRepository.GetUserByEmail(userInfo.Email)
		if err != nil {
			return initiate.INNER_ERROR
		}
		if user != nil {
			return errorx.NewError(400, "邮箱已存在")
		}
	} else {
		if len(userInfo.Phone) == 0 || len(userInfo.Password) == 0 {
			return errorx.NewError(400, "手机号或密码不能为空")
		}
		user, err := u.userRepository.GetUserByPhone(userInfo.Phone)
		if err != nil {
			return initiate.INNER_ERROR
		}
		if user != nil {
			return errorx.NewError(400, "手机号已存在")
		}
	}

	pwd, err := utils.GetPwd(userInfo.Password)
	if err != nil {
		return initiate.INNER_ERROR
	}
	user := &domain.User{
		Base: domain.Base{
			Id:         uuid.NewV4().String(),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
		Phone:    userInfo.Phone,
		Password: pwd,
	}
	if err := u.userRepository.CreateUser(user); err != nil {
		return errorx.NewError(500, "内部错误")
	}
	return nil
}

func (u *UserService) Login(info response.LoginInfo) (string, error) {
	if len(info.LoginType) == 0 {
		return "", initiate.INVALID_ARGUMENT
	}
	var user *domain.User
	var err error
	if info.LoginType == "email" {
		if len(info.Email) == 0 || len(info.Password) == 0 {
			return "", errorx.NewError(400, "邮箱或密码不能为空")
		}
		user, err = u.userRepository.GetUserByEmail(info.Email)
	} else {
		if len(info.Phone) == 0 || len(info.Password) == 0 {
			return "", errorx.NewError(400, "手机号或密码不能为空")
		}
		user, err = u.userRepository.GetUserByPhone(info.Phone)
	}
	if err != nil {
		return "", initiate.INNER_ERROR
	}
	if user == nil {
		return "", errorx.NewError(400, "用户不存在")
	}
	token, err := utils.CreateAccessToken(user)
	if err != nil {
		return "", initiate.INNER_ERROR
	}
	return token, nil
}

func (u *UserService) UpdateUserInfo(userInfo response.RegisterInfo, school *domain.School) error {
	if len(userInfo.RegisterType) == 0 {
		return initiate.INVALID_ARGUMENT
	}
	if len(userInfo.SchoolCode) == 0 {
		return errorx.NewError(400, "学校必须选择!")
	}
	var user *domain.User
	var err error
	if userInfo.RegisterType == "email" {
		if len(userInfo.Email) == 0 {
			return initiate.INNER_ERROR
		}
		user, err = u.userRepository.GetUserByEmail(userInfo.Email)

	} else {
		if len(userInfo.Phone) == 0 {
			return initiate.INNER_ERROR
		}
		user, err = u.userRepository.GetUserByPhone(userInfo.Phone)
	}

	if err != nil {
		return initiate.INNER_ERROR
	}
	if user == nil {
		return errorx.NewError(400, "用户不存在")
	}
	user.SchoolId = school.Id
	user.Username = userInfo.Username
	user.UpdateTime = time.Now()
	user.Avatar = userInfo.Avatar
	user.Gender = userInfo.Gender
	user.Birthday = utils.ParseTime(userInfo.Birthday)
	user.Email = userInfo.Email
	user.Phone = userInfo.Phone
	if err := u.userRepository.UpdateUserInfo(user); err != nil {
		return initiate.INNER_ERROR
	}
	return nil
}
