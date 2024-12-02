package dto

import "atom-go/pkg/datetime"

// 保存用户请求体
type SaveUserRequest struct {
	UserId        int               `json:"userId"`
	UserType      string            `json:"userType"`
	Username      string            `json:"username"`
	Password      string            `json:"password"`
	Nickname      string            `json:"nickname"`
	Avatar        string            `json:"avatar"`
	Email         string            `json:"email"`
	Phone         string            `json:"phone"`
	Gender        string            `json:"gender"`
	LoginLastIp   string            `json:"loginLastIp"`
	LoginLastTime datetime.Datetime `json:"loginLastTime"`
	Status        string            `json:"status"`
	CreateBy      string            `json:"createBy"`
	UpdateBy      string            `json:"updateBy"`
}

// 获取用户列表请求体
type GetUserPageRequest struct {
	PageRequest
	UserType string `query:"userType" form:"userType"`
	Username string `query:"username" form:"username"`
	Nickname string `query:"nickname" form:"nickname"`
	Status   string `query:"status" form:"status"`
}

// 用户id请求体
type UserIdRequest struct {
	UserId int `query:"userId" json:"userId" form:"userId"`
}

// 用户绑定角色请求体
type BindUserRoleRequest struct {
	UserId  int   `json:"userId"`
	RoleIds []int `json:"roleIds"`
}
