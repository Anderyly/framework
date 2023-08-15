package model

import "framework/dal"

type Demo struct {
	dal.Model
	Content string `gorm:"not null" json:"content"` // 内容
}
