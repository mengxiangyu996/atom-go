package router

import (
	"atom-go/internal/controller/admin"
	"atom-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRegister(server *gin.Engine) {

	api := server.Group("/api/admin")
	{
		api.GET("/auth/captcha", (&admin.AuthController{}).Captcha) // 验证码
		api.POST("/auth/login", (&admin.AuthController{}).Login)    // 登录
	}

	// 加入认证中间件（查询）
	api = server.Group("/api/admin", middleware.AdminAuthMiddleware())
	{
		api.GET("/auth/getUserInfo", (&admin.AuthController{}).GetUserInfo)   // 获取授权用户信息
		api.GET("/auth/getUserMenus", (&admin.AuthController{}).GetUserMenus) // 获取授权用户菜单权限

		api.POST("/upload/uploadFile", (&admin.UploadController{}).UploadFile)   // 上传文件
		api.POST("/upload/uploadImage", (&admin.UploadController{}).UploadImage) // 上传图片

		api.GET("/loginLog/getLoginLogPage", (&admin.LoginLogController{}).GetLoginLogPage) // 获取登录日志分页

		api.GET("/operationLog/getOperationLogPage", (&admin.OperationLogController{}).GetOperationLogPage) // 获取操作日志分页
		api.GET("/operationLog/getOperationLogInfo", (&admin.OperationLogController{}).GetOperationLogInfo) // 获取操作日志详情

		api.GET("/user/getUserPage", (&admin.UserController{}).GetUserPage) // 获取用户分页
		api.GET("/user/getUserInfo", (&admin.UserController{}).GetUserInfo) // 获取用户详情

		api.GET("/role/getRolePage", (&admin.RoleController{}).GetRolePage)                   // 获取角色分页
		api.GET("/role/getRoleInfo", (&admin.RoleController{}).GetRoleInfo)                   // 获取角色详情
		api.GET("/role/getRoleBoundMenuTree", (&admin.RoleController{}).GetRoleBoundMenuTree) // 获取角色已绑定的菜单权限树形

		api.GET("/menu/getMenuTree", (&admin.MenuController{}).GetMenuTree) // 获取菜单权限树形
		api.GET("/menu/getMenuInfo", (&admin.MenuController{}).GetMenuInfo) // 获取菜单权限详情
	}

	// 加入操作日志中间件（创建、修改、删除）
	api = server.Group("/api/admin", middleware.AdminAuthMiddleware(), middleware.AdminOperationLogMiddleware())
	{
		api.POST("/auth/updateInfo", (&admin.AuthController{}).UpdateInfo) // 更新个人信息

		api.POST("/user/createUser", (&admin.UserController{}).CreateUser)     // 创建用户
		api.POST("/user/updateUser", (&admin.UserController{}).UpdateUser)     // 更新用户
		api.POST("/user/deleteUser", (&admin.UserController{}).DeleteUser)     // 删除用户
		api.POST("/user/bindUserRole", (&admin.UserController{}).BindUserRole) // 用户绑定角色

		api.POST("/role/createRole", (&admin.RoleController{}).CreateRole)     // 创建角色
		api.POST("/role/updateRole", (&admin.RoleController{}).UpdateRole)     // 更新角色
		api.POST("/role/deleteRole", (&admin.RoleController{}).DeleteRole)     // 删除角色
		api.POST("/role/bindRoleMenu", (&admin.RoleController{}).BindRoleMenu) // 角色绑定菜单权限

		api.POST("/menu/createMenu", (&admin.MenuController{}).CreateMenu) // 创建菜单权限
		api.POST("/menu/updateMenu", (&admin.MenuController{}).UpdateMenu) // 更新菜单权限
		api.POST("/menu/deleteMenu", (&admin.MenuController{}).DeleteMenu) // 删除菜单权限
	}
}
