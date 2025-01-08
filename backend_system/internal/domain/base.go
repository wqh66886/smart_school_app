package domain

import "time"

/**
* description:
* author: wqh
* date: 2025/1/8
 */

type Base struct {
	Id         string    `json:"id" gorm:"primaryKey"`
	CreateTime time.Time `json:"create_time" gorm:"default null"`
	UpdateTime time.Time `json:"update_time" gorm:"default null"`
	Remark     string    `json:"remark" gorm:"type:text;default null"`
}
