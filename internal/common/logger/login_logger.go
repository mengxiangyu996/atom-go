package logger

import (
	ipaddress "atom-go/internal/common/ip-address"
	"atom-go/internal/dto"
	"atom-go/internal/service"
	"atom-go/pkg/datetime"

	"github.com/mileusna/useragent"
)

// 登录日志
type LoginLogger struct {
	Username  string
	Ip        string
	UserAgent string
	Status    string
	Message   string
	LoginTime datetime.Datetime
}

func (l *LoginLogger) Insert() error {

	ipAddress := ipaddress.GetAddress(l.Ip)

	userAgent := useragent.Parse(l.UserAgent)

	return (&service.LoginLogService{}).CreateLoginLog(&dto.CreateLoginLogRequest{
		LoginUsername: l.Username,
		LoginIp:       ipAddress.Ip,
		LoginLocation: ipAddress.Addr,
		Browser:       userAgent.Name,
		Os:            userAgent.OS,
		Status:        l.Status,
		Message:       l.Message,
		LoginTime:     l.LoginTime,
	})
}
