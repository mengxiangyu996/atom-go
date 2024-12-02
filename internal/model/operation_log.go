package model

import "atom-go/pkg/datetime"

// 操作日志模型
type OperationLog struct {
	OperationLogId    int
	Path              string
	Method            string
	Param             string
	Result            string
	OperationIp       string
	OperationLocation string
	OperationName     string
	OperationTime     datetime.Datetime
}
