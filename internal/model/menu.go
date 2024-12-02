package model

import (
	"atom-go/pkg/datetime"

	"gorm.io/gorm"
)

// 菜单权限模型
type Menu struct {
	MenuId     int
	MenuName   string
	ParentId   int `gorm:"default:0"`
	MenuType   string
	Sort       int `gorm:"default:0"`
	Component  string
	Icon       string
	Path       string
	Method     string
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
}
