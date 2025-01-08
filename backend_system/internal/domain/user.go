package domain

import "time"

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type User struct {
	Base
	Phone    string    `json:"phone" gorm:"unique;not null"`
	Username string    `json:"username" gorm:"not null"`
	Password string    `json:"password" gorm:"not null"`
	Gender   string    `json:"gender"`
	Birthday time.Time `json:"birthday"`
	Avatar   string    `json:"avatar" gorm:"default null"`
	SchoolId string    `json:"school_id" gorm:"not null"`
}

func (u *User) TableName() string {
	return "user"
}
