package dto

import "atom-go/pkg/datetime"

// 用户Token响应体
type UserTokenResponse struct {
	UserId   int    `json:"userId"`
	UserType string `json:"userType"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Status   string `json:"status"`
}

// 用户列表响应体
type UserListResponse struct {
	UserId        int               `json:"userId"`
	UserType      string            `json:"userType"`
	Username      string            `json:"username"`
	Nickname      string            `json:"nickname"`
	Avatar        string            `json:"avatar"`
	LoginLastIp   string            `json:"loginLastIp"`
	LoginLastTime datetime.Datetime `json:"loginLastTime"`
	Status        string            `json:"status"`
	CreateBy      string            `json:"createBy"`
	CreateTime    datetime.Datetime `json:"createTime"`
}

// 用户详情响应体
type UserInfoResponse struct {
	UserId        int               `json:"userId"`
	UserType      string            `json:"userType"`
	Username      string            `json:"username"`
	Nickname      string            `json:"nickname"`
	Avatar        string            `json:"avatar"`
	Email         string            `json:"email"`
	Phone         string            `json:"phone"`
	Gender        string            `json:"gender"`
	LoginLastIp   string            `json:"loginLastIp"`
	LoginLastTime datetime.Datetime `json:"loginLastTime"`
	Status        string            `json:"status"`
	CreateBy      string            `json:"createBy"`
	CreateTime    datetime.Datetime `json:"createTime"`
	UpdateBy      string            `json:"updateBy"`
	UpdateTime    datetime.Datetime `json:"updateTime"`
}
