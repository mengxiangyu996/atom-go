package dto

import "atom-go/pkg/datetime"

// 登录日志列表响应体
type LoginLogListResponse struct {
	LoginLogId    int               `json:"loginLogId"`
	LoginUsername string            `json:"loginUsername"`
	LoginIp       string            `json:"loginIp"`
	LoginLocation string            `json:"loginLocation"`
	Browser       string            `json:"browser"`
	Os            string            `json:"os"`
	Status        string            `json:"status"`
	Message       string            `json:"message"`
	LoginTime     datetime.Datetime `json:"loginTime"`
}
