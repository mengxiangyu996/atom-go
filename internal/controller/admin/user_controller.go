package admin

import (
	"atom-go/internal/common/password"
	"atom-go/internal/dto"
	"atom-go/internal/service"
	"atom-go/internal/validator"
	"atom-go/pkg/response"

	"github.com/gin-gonic/gin"
)

// 用户控制器
type UserController struct{}

// 获取用户分页
func (*UserController) GetUserPage(ctx *gin.Context) {

	var param dto.GetUserPageRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	users, total := (&service.UserService{}).GetUserPage(&param)

	response.NewSuccess().SetData(dto.PageResponse{
		List:  users,
		Total: total,
	}).Send(ctx)
}

// 获取用户详情
func (*UserController) GetUserInfo(ctx *gin.Context) {

	var param dto.UserIdRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	user := (&service.UserService{}).GetUserInfoByUserId(param.UserId)

	response.NewSuccess().SetData(user).Send(ctx)
}

// 创建用户
func (*UserController) CreateUser(ctx *gin.Context) {

	var param dto.SaveUserRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if err := validator.CreateUserValidator(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if user := (&service.UserService{}).GetUserInfoByUsername(param.Username); user.UserId > 0 {
		response.NewError().SetMessage("用户名已存在").Send(ctx)
		return
	}

	param.Password = password.Generate(param.Password)

	param.CreateBy = ctx.GetString("nickname")

	if err := (&service.UserService{}).CreateUser(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}

// 更新用户
func (*UserController) UpdateUser(ctx *gin.Context) {

	var param dto.SaveUserRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if err := validator.UpdateUserValidator(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if param.UserId == 1 && param.UserType != "SUPER_ADMIN" {
		response.NewError().SetMessage("无法修改超级管理员的角色").Send(ctx)
		return
	}

	// 排除账号，只能自己修改自己的账号
	param.Username = ""

	// 排除密码，只能自己修改自己的密码
	param.Password = ""

	param.UpdateBy = ctx.GetString("nickname")

	if err := (&service.UserService{}).UpdateUserByUserId(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}

// 删除用户
func (*UserController) DeleteUser(ctx *gin.Context) {

	var param dto.UserIdRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if param.UserId == 1 {
		response.NewError().SetMessage("无法删除超级管理员").Send(ctx)
		return
	}

	if err := (&service.UserService{}).DeleteUserByUserId(param.UserId); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}

// 用户绑定角色
func (*UserController) BindUserRole(ctx *gin.Context) {

	var param dto.BindUserRoleRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if err := (&service.UserRoleService{}).BindUserRole(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}
