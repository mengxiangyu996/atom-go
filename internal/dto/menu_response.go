package dto

import "atom-go/pkg/datetime"

// 菜单权限列表响应体
type MenuListResponse struct {
	MenuId     int               `json:"menuId"`
	MenuName   string            `json:"menuName"`
	ParentId   int               `json:"parentId"`
	MenuType   string            `json:"menuType"`
	Sort       int               `json:"sort"`
	Component  string            `json:"component"`
	Icon       string            `json:"icon"`
	Path       string            `json:"path"`
	Method     string            `json:"method"`
	Status     string            `json:"status"`
	CreateBy   string            `json:"createBy"`
	CreateTime datetime.Datetime `json:"createTime"`
}

// 菜单权限详情响应体
type MenuInfoResponse struct {
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
}

// 菜单权限树形响应体
type MenuTreeResponse struct {
	MenuListResponse
	Children []MenuTreeResponse `json:"children"`
}
