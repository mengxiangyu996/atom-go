package admin

import (
	"atom-go/internal/common/captcha"
	"atom-go/internal/common/logger"
	"atom-go/internal/common/password"
	"atom-go/internal/dto"
	"atom-go/internal/service"
	"atom-go/internal/token"
	"atom-go/internal/validator"
	"atom-go/pkg/datetime"
	"atom-go/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
)

// 授权控制器
type AuthController struct{}

// 验证码
func (*AuthController) Captcha(ctx *gin.Context) {

	captcha := captcha.NewCaptcha()

	id, b64s := captcha.Generate()

	response.NewSuccess().SetData(dto.CaptchaResponse{
		CaptchaId:    id,
		CaptchaImage: b64s,
	}).Send(ctx)
}

// 登录
func (*AuthController) Login(ctx *gin.Context) {

	var param dto.LoginRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if err := validator.LoginValidator(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		go (&logger.LoginLogger{
			Username:  param.Username,
			Ip:        ctx.ClientIP(),
			UserAgent: ctx.Request.UserAgent(),
			Status:    "1",
			Message:   err.Error(),
			LoginTime: datetime.Datetime{Time: time.Now()},
		}).Insert()
		return
	}

	captcha := captcha.NewCaptcha()
	if !captcha.Verify(param.CaptchaId, param.CaptchaCode) {
		response.NewError().SetMessage("验证码错误").Send(ctx)
		go (&logger.LoginLogger{
			Username:  param.Username,
			Ip:        ctx.ClientIP(),
			UserAgent: ctx.Request.UserAgent(),
			Status:    "1",
			Message:   "验证码错误",
			LoginTime: datetime.Datetime{Time: time.Now()},
		}).Insert()
		return
	}

	user := (&service.UserService{}).GetUserInfoByUsername(param.Username)
	if user.UserId <= 0 || user.Status != "0" {
		response.NewError().SetMessage("用户不存在或被禁用").Send(ctx)
		go (&logger.LoginLogger{
			Username:  param.Username,
			Ip:        ctx.ClientIP(),
			UserAgent: ctx.Request.UserAgent(),
			Status:    "1",
			Message:   "用户不存在或被禁用",
			LoginTime: datetime.Datetime{Time: time.Now()},
		}).Insert()
		return
	}

	// 更新登录ip和时间
	(&service.UserService{}).UpdateUserByUserId(&dto.SaveUserRequest{
		UserId:        user.UserId,
		LoginLastIp:   ctx.ClientIP(),
		LoginLastTime: datetime.Datetime{Time: time.Now()},
	})

	token := token.GetAdminClaims(user).GenerateToken()

	go (&logger.LoginLogger{
		Username:  param.Username,
		Ip:        ctx.ClientIP(),
		UserAgent: ctx.Request.UserAgent(),
		Status:    "0",
		Message:   "登录成功",
		LoginTime: datetime.Datetime{Time: time.Now()},
	}).Insert()

	response.NewSuccess().SetData(token).Send(ctx)
}

// 获取授权用户信息
func (*AuthController) GetUserInfo(ctx *gin.Context) {

	userId := ctx.GetInt("userId")

	user := (&service.UserService{}).GetUserInfoByUserId(userId)

	response.NewSuccess().SetData(user).Send(ctx)
}

// 获取授权用户菜单权限
func (*AuthController) GetUserMenus(ctx *gin.Context) {

	adminClaims, _ := token.ParseAdminToken(ctx)

	menus := make([]dto.MenuListResponse, 0)

	if adminClaims.UserType == "SUPER_ADMIN" {
		// 超级管理员拥有所有权限
		// 只需要目录和菜单，以便前端渲染
		menus = (&service.MenuService{}).GetMenuList(&dto.GetMenuListRequest{
			Status:   "0",
			MenuType: []string{"M", "C"},
		})
	} else {
		// 非超级管理员，根据角色查询菜单权限
		roleIds := (&service.UserRoleService{}).GetRoleIdsByUserIds([]int{adminClaims.UserId})
		menuIds := (&service.RoleMenuService{}).GetMenuIdsByRoleIds(roleIds)
		// 只需要目录和菜单，以便前端渲染
		menus = (&service.MenuService{}).GetMenuList(&dto.GetMenuListRequest{
			MenuIds:  menuIds,
			Status:   "0",
			MenuType: []string{"M", "C"},
		})
	}

	// 菜单权限列表转为树形结构
	tree := (&service.MenuService{}).MenuListToTree(menus, 0)

	response.NewSuccess().SetData(tree).Send(ctx)
}

// 更新个人信息
func (*AuthController) UpdateInfo(ctx *gin.Context) {

	var param dto.SaveUserRequest

	if err := ctx.ShouldBindJSON(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	if err := validator.UpdateInfoValidator(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	param.UserId = ctx.GetInt("userId")
	param.UpdateBy = ctx.GetString("nickname")

	if param.Username != "" {
		if user := (&service.UserService{}).GetUserInfoByUsername(param.Username); user.UserId > 0 && user.UserId != param.UserId {
			response.NewError().SetMessage("用户名已存在").Send(ctx)
			return
		}
	}

	if param.Password != "" {
		param.Password = password.Generate(param.Password)
	}

	// 无法修改自己的角色
	param.UserType = ""

	if err := (&service.UserService{}).UpdateUserByUserId(&param); err != nil {
		response.NewError().SetMessage(err.Error()).Send(ctx)
		return
	}

	response.NewSuccess().Send(ctx)
}
