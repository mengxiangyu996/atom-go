package dto

import "atom-go/pkg/datetime"

// 创建登录日志请求体
type CreateLoginLogRequest struct {
	LoginUsername string            `json:"loginUsername"`
	LoginIp       string            `json:"loginIp"`
	LoginLocation string            `json:"loginLocation"`
	Browser       string            `json:"browser"`
	Os            string            `json:"os"`
	Status        string            `json:"status"`
	Message       string            `json:"message"`
	LoginTime     datetime.Datetime `json:"loginTime"`
}

// 获取登录日志列表请求体
type GetLoginLogPageRequest struct {
	PageRequest
	LoginUsername string   `query:"loginUsername" form:"loginUsername"`
	LoginLocation string   `query:"loginLocation" form:"loginLocation"`
	Status        string   `query:"status" form:"status"`
	LoginTime     []string `query:"loginTime" form:"loginTime"`
}
