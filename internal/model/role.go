package model

import (
	"atom-go/pkg/datetime"

	"gorm.io/gorm"
)

// 角色模型
type Role struct {
	RoleId     int
	RoleName   string
	Status     string `gorm:"default:0"`
	CreateBy   string
	CreateTime datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy   string
	UpdateTime datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime gorm.DeletedAt
}
