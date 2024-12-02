package dto

import "atom-go/pkg/datetime"

// 角色列表响应体
type RoleListResponse struct {
	RoleId     int               `json:"roleId"`
	RoleName   string            `json:"roleName"`
	Status     string            `json:"status"`
	CreateBy   string            `json:"createBy"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// 角色详情响应体
type RoleInfoResponse struct {
	RoleId     int               `json:"roleId"`
	RoleName   string            `json:"roleName"`
	Status     string            `json:"status"`
	CreateBy   string            `json:"createBy"`
	CreateTime datetime.Datetime `json:"createTime"`
	UpdateBy   string            `json:"updateBy"`
	UpdateTime datetime.Datetime `json:"updateTime"`
}
