package admin

import (
	"atom-go/internal/dto"
	"atom-go/internal/service"
	"atom-go/pkg/response"

	"github.com/gin-gonic/gin"
)

// 登录日志控制器
type LoginLogController struct{}

// 获取登录日志分页
func (*LoginLogController) GetLoginLogPage(ctx *gin.Context) {

	var param dto.GetLoginLogPageRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	loginLogs, total := (&service.LoginLogService{}).GetLoginLogPage(&param)

	response.NewSuccess().SetData(dto.PageResponse{
		List:  loginLogs,
		Total: total,
	}).Json(ctx)
}
