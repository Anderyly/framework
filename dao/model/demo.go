package model

import "gorm.io/gorm"

type Demo struct {
	gorm.Model
	Content    string `gorm:"not null"` // 内容
	CreateTime int64  `gorm:"not null"` // 创建时间
	UpdateTime int64  `gorm:"not null"` // 更新时间
}
