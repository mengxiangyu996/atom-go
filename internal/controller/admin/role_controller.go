package admin

import (
	"atom-go/internal/dto"
	"atom-go/internal/service"
	"atom-go/internal/validator"
	"atom-go/pkg/response"

	"github.com/gin-gonic/gin"
)

// 角色控制器
type RoleController struct{}

// 获取角色分页
func (*RoleController) GetRolePage(ctx *gin.Context) {

	var param dto.GetRolePageRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	roles, total := (&service.RoleService{}).GetRoleList(&param)

	response.NewSuccess().SetData(dto.PageResponse{
		List:  roles,
		Total: total,
	}).Send(ctx)
}

// 获取角色详情
func (*RoleController) GetRoleInfo(ctx *gin.Context) {

	var param dto.RoleIdRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	role := (&service.RoleService{}).GetRoleInfoByRoleId(param.RoleId)

	response.NewSuccess().SetData(role).Send(ctx)
}

// 创建角色
func (*RoleController) CreateRole(ctx *gin.Context) {

	var param dto.SaveRoleRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if err := validator.CreateRoleValidator(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if role := (&service.RoleService{}).GetRoleInfoByRoleName(param.RoleName); role.RoleId > 0 {
		response.NewError().SetMessage("角色名称已存在").Send(ctx)
		return
	}

	param.CreateBy = ctx.GetString("nickname")

	if err := (&service.RoleService{}).CreateRole(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}

// 更新角色
func (*RoleController) UpdateRole(ctx *gin.Context) {

	var param dto.SaveRoleRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if err := validator.UpdateRoleValidator(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if role := (&service.RoleService{}).GetRoleInfoByRoleName(param.RoleName); role.RoleId > 0 && role.RoleId != param.RoleId {
		response.NewError().SetMessage("角色名称已存在").Send(ctx)
		return
	}

	param.UpdateBy = ctx.GetString("nickname")

	if err := (&service.RoleService{}).UpdateRoleByRoleId(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}

// 删除角色
func (*RoleController) DeleteRole(ctx *gin.Context) {

	var param dto.RoleIdRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if userIds := (&service.UserRoleService{}).GetUserIdsByRoleIds([]int{param.RoleId}); len(userIds) > 0 {
		response.NewError().SetMessage("请先删除角色关联的用户").Send(ctx)
		return
	}

	if err := (&service.RoleService{}).DeleteRoleByRoleId(param.RoleId); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}

// 角色绑定菜单权限
func (*RoleController) BindRoleMenu(ctx *gin.Context) {

	var param dto.BindRoleMenuRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if err := (&service.RoleMenuService{}).BindRoleMenu(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}

// 获取角色已绑定的菜单权限树形
func (*RoleController) GetRoleBoundMenuTree(ctx *gin.Context) {

	var param dto.RoleIdRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	menuIds := (&service.RoleMenuService{}).GetMenuIdsByRoleIds([]int{param.RoleId})

	menus := (&service.MenuService{}).GetMenuList(&dto.GetMenuListRequest{
		MenuIds: menuIds,
		Status:  "0",
	})

	tree := (&service.MenuService{}).MenuListToTree(menus, 0)

	response.NewSuccess().SetData(tree).Send(ctx)
}
