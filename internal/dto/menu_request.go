package dto

// 保存菜单权限请求体
type SaveMenuRequest struct {
	MenuId    int    `json:"menuId"`
	MenuName  string `json:"menuName"`
	ParentId  int    `json:"parentId"`
	MenuType  string `json:"menuType"`
	Sort      int    `json:"sort"`
	Component string `json:"component"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	Method    string `json:"method"`
	Status    string `json:"status"`
	CreateBy  string `json:"createBy"`
	UpdateBy  string `json:"updateBy"`
}

// 获取菜单权限列表请求体
type GetMenuListRequest struct {
	MenuIds  []int    `query:"menuIds" form:"menuIds"`
	MenuName string   `query:"menuName" form:"menuName"`
	MenuType []string `query:"menuType" form:"menuType"`
	Status   string   `query:"status" form:"status"`
}

// 菜单权限id请求体
type MenuIdRequest struct {
	MenuId int `query:"menuId" json:"menuId" form:"menuId"`
}
