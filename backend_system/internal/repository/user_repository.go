package mapper

import "github.com/wqh/smart/school/system/internal/domain"

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type UserMapper interface {
	CreateUser(user *domain.User) error
	GetUserByUsername(username string) (*domain.User, error)
}
