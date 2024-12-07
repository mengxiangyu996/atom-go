package middleware

import (
	"atom-go/internal/common/logger"
	"atom-go/internal/service"
	"atom-go/internal/token"
	"atom-go/pkg/datetime"
	"atom-go/pkg/response"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

// 后台鉴权中间件
func AdminAuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		adminClaims, err := token.ParseAdminToken(ctx)
		if err != nil {
			response.NewError().SetCode(10401).SetMessage(err.Error()).Json(ctx)
			ctx.Abort()
			return
		}

		if adminClaims.Subject != "admin" {
			response.NewError().SetCode(10401).SetMessage("用户类型错误").Json(ctx)
			ctx.Abort()
			return
		}

		ctx.Set("userId", adminClaims.UserId)
		ctx.Set("nickname", adminClaims.Nickname)

		// 超级管理员跳过后续验证
		if adminClaims.UserId == 1 || adminClaims.UserType == "SUPER_ADMIN" {
			ctx.Next()
			return
		}

		if adminClaims.Status != "0" {
			response.NewError().SetCode(10401).SetMessage("用户被禁用").Json(ctx)
			ctx.Abort()
			return
		}

		menu := (&service.MenuService{}).GetMenuByPathAndMethod(ctx.Request.URL.Path, ctx.Request.Method)
		if menu.MenuId <= 0 {
			response.NewError().SetCode(10401).SetMessage("请求权限不存在").Json(ctx)
			ctx.Abort()
			return
		}

		// 鉴权逻辑
		roleIds := (&service.UserRoleService{}).GetRoleIdsByUserIds([]int{adminClaims.UserId})
		if len(roleIds) == 0 {
			response.NewError().SetCode(10401).SetMessage("用户无角色权限").Json(ctx)
			ctx.Abort()
			return
		}

		if !(&service.RoleMenuService{}).IsBindRoleMenu(roleIds, menu.MenuId) {
			response.NewError().SetCode(10401).SetMessage("用户无请求权限").Json(ctx)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// 重写gin的ResponseWriter，用户接收响应体
type responseWriter struct {
	gin.ResponseWriter
	Body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {

	rw.Body.Write(b)

	return rw.ResponseWriter.Write(b)
}

func (rw *responseWriter) WriteString(s string) (int, error) {

	rw.Body.WriteString(s)

	return rw.ResponseWriter.WriteString(s)
}

// 后台操作日志中间件
func AdminOperationLogMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		// 因读取请求体后，请求体的数据流会被消耗完毕，未避免EOF错误，需要缓存请求体，并且每次使用后需要重新赋值给ctx.Request.Body
		bodyBytes, _ := ctx.GetRawData()
		// 将缓存的请求体重新赋值给ctx.Request.Body，供下方ctx.ShouldBind使用
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		rw := &responseWriter{
			ResponseWriter: ctx.Writer,
			Body:           bytes.NewBufferString(""),
		}

		paramData := make(map[string]interface{}, 0)
		ctx.ShouldBind(&paramData)

		// 因ctx.ShouldBind后，请求体的数据流会被消耗完毕，需要将缓存的请求体重新赋值给ctx.Request.Body
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// 将query参数转为map并添加到请求参数中，用query-key的形式以便区分
		for key, value := range ctx.Request.URL.Query() {
			paramData["query-"+key] = value
		}

		param, _ := json.Marshal(&paramData)

		ctx.Writer = rw

		operationLogger := &logger.OperationLogger{
			Path:          ctx.Request.URL.Path,
			Method:        ctx.Request.Method,
			Param:         string(param),
			Ip:            ctx.ClientIP(),
			OperationName: ctx.GetString("nickname"),
			OperationTime: datetime.Datetime{Time: time.Now()},
		}

		ctx.Next()

		operationLogger.Result = rw.Body.String()
		operationLogger.Insert()
	}
}
