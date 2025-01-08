package domain

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type School struct {
	Base
	Name    string `json:"name" gorm:"not null"`
	Code    string `json:"code" gorm:"not null"`
	Logo    string `json:"logo" gorm:"default null"`
	Address string `json:"address" gorm:"not null"`
	Phone   string `json:"phone" gorm:"not null"`
	Email   string `json:"email" gorm:"not null"`
	Website string `json:"website" gorm:"not null"`
}

func (s *School) TableName() string {
	return "school"
}
