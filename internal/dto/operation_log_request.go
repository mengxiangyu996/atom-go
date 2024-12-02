package dto

import "atom-go/pkg/datetime"

// 创建操作日志请求体
type CreateOperationLogRequest struct {
	Path              string            `json:"path"`
	Method            string            `json:"method"`
	Param             string            `json:"param"`
	Result            string            `json:"result"`
	OperationIp       string            `json:"operationIp"`
	OperationLocation string            `json:"operationLocation"`
	OperationName     string            `json:"operationName"`
	OperationTime     datetime.Datetime `json:"operationTime"`
}

// 获取操作日志列表请求体
type GetOperationLogPageRequest struct {
	PageRequest
	Path          string   `query:"path" form:"path"`
	Method        string   `query:"method" form:"method"`
	OperationName string   `query:"operationName" form:"operationName"`
	OperationTime []string `query:"operationTime" form:"operationTime"`
}

// 操作日志id请求体
type OperationLogIdRequest struct {
	OperationLogId int `query:"operationLogId" form:"operationLogId"`
}
