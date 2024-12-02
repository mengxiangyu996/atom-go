package dto

// 分页请求体
type PageRequest struct {
	Page     int `query:"page" form:"page"`
	PageSize int `query:"pageSize" form:"pageSize"`
}
