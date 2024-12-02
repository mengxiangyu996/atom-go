package service

import (
	"atom-go/internal/dto"
	"atom-go/internal/model"
	"atom-go/pkg/dal"
)

// 登录日志服务
type LoginLogService struct{}

// 创建登录日志
func (l *LoginLogService) CreateLoginLog(param *dto.CreateLoginLogRequest) error {

	return dal.Gorm.Model(model.LoginLog{}).Create(&model.LoginLog{
		LoginUsername: param.LoginUsername,
		LoginIp:       param.LoginIp,
		LoginLocation: param.LoginLocation,
		Browser:       param.Browser,
		Os:            param.Os,
		Status:        param.Status,
		Message:       param.Message,
		LoginTime:     param.LoginTime,
	}).Error
}

// 获取登录日志列表
func (l *LoginLogService) GetLoginLogPage(param *dto.GetLoginLogPageRequest) ([]dto.LoginLogListResponse, int) {

	var count int64
	loginLogs := make([]dto.LoginLogListResponse, 0)

	tx := dal.Gorm.Model(model.LoginLog{}).Order("login_log_id DESC")

	if param.LoginUsername != "" {
		tx.Where("login_username = ?", param.LoginUsername)
	}

	if param.LoginLocation != "" {
		tx.Where("login_location = ?", param.LoginLocation)
	}

	if param.Status != "" {
		tx.Where("status = ?", param.Status)
	}

	if len(param.LoginTime) > 0 {
		tx.Where("login_time BETWEEN ? AND ?", param.LoginTime[0], param.LoginTime[1])
	}

	tx.Count(&count).Offset((param.Page - 1) * param.PageSize).Limit(param.PageSize).Find(&loginLogs)

	return loginLogs, int(count)
}
