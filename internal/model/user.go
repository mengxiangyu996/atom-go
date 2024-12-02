package model

import (
	"atom-go/pkg/datetime"

	"gorm.io/gorm"
)

// 用户模型
type User struct {
	UserId        int
	UserType      string `gorm:"default:USER"`
	Username      string
	Password      string
	Nickname      string
	Avatar        string
	Email         string
	Phone         string
	Gender        string `gorm:"default:0"`
	LoginLastIp   string
	LoginLastTime datetime.Datetime `gorm:"default:null"`
	Status        string            `gorm:"default:0"`
	CreateBy      string
	CreateTime    datetime.Datetime `gorm:"autoCreateTime"`
	UpdateBy      string
	UpdateTime    datetime.Datetime `gorm:"autoUpdateTime"`
	DeleteTime    gorm.DeletedAt
}
