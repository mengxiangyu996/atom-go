package dto

// 分页响应体
type PageResponse struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}
