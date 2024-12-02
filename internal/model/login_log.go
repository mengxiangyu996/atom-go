package model

import "atom-go/pkg/datetime"

// 登录日志模型
type LoginLog struct {
	LoginLogId    int
	LoginUsername string
	LoginIp       string
	LoginLocation string
	Browser       string
	Os            string
	Status        string `gorm:"default:0"`
	Message       string
	LoginTime     datetime.Datetime
}
