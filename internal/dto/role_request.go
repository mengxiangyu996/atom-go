package dto

// 保存角色请求体
type SaveRoleRequest struct {
	RoleId   int    `json:"roleId"`
	RoleName string `json:"roleName"`
	Status   string `json:"status"`
	CreateBy string `json:"createBy"`
	UpdateBy string `json:"updateBy"`
}

// 获取角色分页请求体
type GetRolePageRequest struct {
	PageRequest
	RoleName string `query:"roleName" form:"roleName"`
	Status   string `query:"status" form:"status"`
}

// 角色id请求体
type RoleIdRequest struct {
	RoleId int `query:"roleId" json:"roleId" form:"roleId"`
}

// 角色绑定菜单权限请求体
type BindRoleMenuRequest struct {
	RoleId  int   `json:"roleId"`
	MenuIds []int `json:"menuIds"`
}
