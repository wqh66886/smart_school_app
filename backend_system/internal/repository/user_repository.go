package repository

import "github.com/wqh/smart/school/system/internal/domain"

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByPhone(phone string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	UpdateUserInfo(user *domain.User) error
}
