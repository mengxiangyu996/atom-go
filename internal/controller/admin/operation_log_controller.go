package admin

import (
	"atom-go/internal/dto"
	"atom-go/internal/service"
	"atom-go/pkg/response"

	"github.com/gin-gonic/gin"
)

// 操作日志控制器
type OperationLogController struct{}

// 获取操作日志分页
func (*OperationLogController) GetOperationLogPage(ctx *gin.Context) {

	var param dto.GetOperationLogPageRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	operationLogs, total := (&service.OperationLogService{}).GetOperationLogPage(&param)

	response.NewSuccess().SetPageData(operationLogs, total).Json(ctx)
}

// 获取操作日志详情
func (*OperationLogController) GetOperationLogInfo(ctx *gin.Context) {

	var param dto.OperationLogIdRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	operationLog := (&service.OperationLogService{}).GetOperationLogInfo(param.OperationLogId)

	response.NewSuccess().SetData("data", operationLog).Json(ctx)
}
