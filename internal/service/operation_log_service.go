package service

import (
	"atom-go/internal/dto"
	"atom-go/internal/model"
	"atom-go/pkg/dal"
)

// 操作日志服务
type OperationLogService struct{}

// 创建操作日志
func (o *OperationLogService) CreateOperationLog(param *dto.CreateOperationLogRequest) error {

	return dal.Gorm.Model(model.OperationLog{}).Create(&model.OperationLog{
		Path:              param.Path,
		Method:            param.Method,
		Param:             param.Param,
		Result:            param.Result,
		OperationIp:       param.OperationIp,
		OperationLocation: param.OperationLocation,
		OperationName:     param.OperationName,
		OperationTime:     param.OperationTime,
	}).Error
}

// 获取操作日志分页
func (o *OperationLogService) GetOperationLogPage(param *dto.GetOperationLogPageRequest) ([]dto.OperationLogListResponse, int) {

	var count int64
	operationLogs := make([]dto.OperationLogListResponse, 0)

	tx := dal.Gorm.Model(model.OperationLog{}).Order("operation_log_id DESC")

	if param.Path != "" {
		tx.Where("path = ?", param.Path)
	}

	if param.Method != "" {
		tx.Where("method = ?", param.Method)
	}

	if param.OperationName != "" {
		tx.Where("operation_name LIKE ?", "%"+param.OperationName+"%")
	}

	if len(param.OperationTime) > 0 {
		tx.Where("operation_time BETWEEN ? AND ?", param.OperationTime[0], param.OperationTime[1])
	}

	tx.Count(&count).Offset((param.Page - 1) * param.PageSize).Limit(param.PageSize).Find(&operationLogs)

	return operationLogs, int(count)
}

// 获取操作日志详情
func (o *OperationLogService) GetOperationLogInfo(operationLogId int) dto.OperationLogInfoResponse {

	var operationLog dto.OperationLogInfoResponse

	dal.Gorm.Model(model.OperationLog{}).Where("operation_log_id = ?", operationLogId).First(&operationLog)

	return operationLog
}
