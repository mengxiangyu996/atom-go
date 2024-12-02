package service

import (
	"atom-go/internal/dto"
	"atom-go/internal/model"
	"atom-go/pkg/dal"
)

// 角色菜单权限服务
type RoleMenuService struct{}

// 角色绑定菜单权限
func (r *RoleMenuService) BindRoleMenu(param *dto.BindRoleMenuRequest) error {

	tx := dal.Gorm.Begin()

	// 删除之前所绑定的菜单权限
	if err := tx.Model(model.RoleMenu{}).Where("role_id = ?", param.RoleId).Delete(&model.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 绑定新菜单权限
	for _, menuId := range param.MenuIds {
		if err := tx.Model(model.RoleMenu{}).Create(&model.RoleMenu{RoleId: param.RoleId, MenuId: menuId}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// 根据角色id集合查询菜单权限id集合
func (r *RoleMenuService) GetMenuIdsByRoleIds(roleIds []int) []int {

	var menuIds []int

	dal.Gorm.Model(model.RoleMenu{}).Where("role_id IN ?", roleIds).Pluck("menu_id", &menuIds)

	return menuIds
}

// 根据角色id和菜单id查询是否绑定
func (r *RoleMenuService) IsBindRoleMenu(roleIds []int, menuId int) bool {

	var roleMenu model.RoleMenu

	return dal.Gorm.Model(model.RoleMenu{}).Where("role_id IN ? AND menu_id = ?", roleIds, menuId).First(&roleMenu).RowsAffected > 0
}
