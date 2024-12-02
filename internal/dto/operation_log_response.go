package dto

import "atom-go/pkg/datetime"

// 操作日志列表响应体
type OperationLogListResponse struct {
	OperationLogId    int               `json:"operationLogId"`
	Path              string            `json:"path"`
	Method            string            `json:"method"`
	OperationIp       string            `json:"operationIp"`
	OperationLocation string            `json:"operationLocation"`
	OperationName     string            `json:"operationName"`
	OperationTime     datetime.Datetime `json:"operationTime"`
}

// 操作日志详情响应体
type OperationLogInfoResponse struct {
	OperationLogId    int               `json:"operationLogId"`
	Path              string            `json:"path"`
	Method            string            `json:"method"`
	Param             string            `json:"param"`
	Result            string            `json:"result"`
	OperationIp       string            `json:"operationIp"`
	OperationLocation string            `json:"operationLocation"`
	OperationName     string            `json:"operationName"`
	OperationTime     datetime.Datetime `json:"operationTime"`
}
