package usecase

import (
	"github.com/wqh/smart/school/system/internal/domain"
	"gorm.io/gorm"
)

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type userUseCase struct {
	db *gorm.DB
}

func NewUserUseCase(db *gorm.DB) *userUseCase {
	return &userUseCase{db: db}
}

func (u *userUseCase) CreateUser(user *domain.User) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(user).Error
	})
	return err
}

func (u *userUseCase) GetUserByPhone(phone string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Model(&domain.User{}).Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userUseCase) UpdateUserInfo(user *domain.User) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		return tx.Save(user).Error
	})
	return err
}
