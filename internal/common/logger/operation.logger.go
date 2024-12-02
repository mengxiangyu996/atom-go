package logger

import (
	ipaddress "atom-go/internal/common/ip-address"
	"atom-go/internal/dto"
	"atom-go/internal/service"
	"atom-go/pkg/datetime"
)

// 操作日志
type OperationLogger struct {
	Path          string
	Method        string
	Param         string
	Result        string
	Ip            string
	OperationName string
	OperationTime datetime.Datetime
}

func (o *OperationLogger) Insert() error {

	ipAddress := ipaddress.GetAddress(o.Ip)

	return (&service.OperationLogService{}).CreateOperationLog(&dto.CreateOperationLogRequest{
		Path:              o.Path,
		Method:            o.Method,
		Param:             o.Param,
		Result:            o.Result,
		OperationIp:       ipAddress.Ip,
		OperationLocation: ipAddress.Addr,
		OperationName:     o.OperationName,
		OperationTime:     o.OperationTime,
	})
}
