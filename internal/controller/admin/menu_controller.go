package admin

import (
	"atom-go/internal/dto"
	"atom-go/internal/service"
	"atom-go/internal/validator"
	"atom-go/pkg/response"

	"github.com/gin-gonic/gin"
)

// 菜单权限控制器
type MenuController struct{}

// 获取菜单权限树形
func (*MenuController) GetMenuTree(ctx *gin.Context) {

	var param dto.GetMenuListRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	menus := (&service.MenuService{}).GetMenuList(&param)

	tree := (&service.MenuService{}).MenuListToTree(menus, 0)

	response.NewSuccess().SetData(tree).Json(ctx)
}

// 获取菜单权限详情
func (*MenuController) GetMenuInfo(ctx *gin.Context) {

	var param dto.MenuIdRequest

	if err := ctx.ShouldBindQuery(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	menu := (&service.MenuService{}).GetMenuInfoByMenuId(param.MenuId)

	response.NewSuccess().SetData(menu).Json(ctx)
}

// 创建菜单权限
func (*MenuController) CreateMenu(ctx *gin.Context) {

	var param dto.SaveMenuRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	if err := validator.CreateMenuValidator(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	if param.Path != "" || param.Method != "" {
		if menu := (&service.MenuService{}).GetMenuByPathAndMethod(param.Path, param.Method); menu.MenuId > 0 {
			response.NewError().SetMessage("菜单权限已存在").Json(ctx)
			return
		}
	}

	param.CreateBy = ctx.GetString("nickname")

	if err := (&service.MenuService{}).CreateMenu(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 更新菜单权限
func (*MenuController) UpdateMenu(ctx *gin.Context) {

	var param dto.SaveMenuRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	if err := validator.UpdateMenuValidator(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	if param.Path != "" || param.Method != "" {
		if menu := (&service.MenuService{}).GetMenuByPathAndMethod(param.Path, param.Method); menu.MenuId > 0 && menu.MenuId != param.MenuId {
			response.NewError().SetMessage("菜单权限已存在").Json(ctx)
			return
		}
	}

	param.UpdateBy = ctx.GetString("nickname")

	if err := (&service.MenuService{}).UpdateMenuByMenuId(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}

// 删除菜单权限
func (*MenuController) DeleteMenu(ctx *gin.Context) {

	var param dto.MenuIdRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	if err := (&service.MenuService{}).DeleteMenuByMenuId(param.MenuId); err != nil {
		response.NewError().SetMessage(err.Error()).Json(ctx)
		return
	}

	response.NewSuccess().Json(ctx)
}
