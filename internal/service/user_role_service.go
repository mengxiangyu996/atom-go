package service

import (
	"atom-go/internal/dto"
	"atom-go/internal/model"
	"atom-go/pkg/dal"
)

// 用户角色服务
type UserRoleService struct{}

// 用户绑定角色
func (u *UserRoleService) BindUserRole(param *dto.BindUserRoleRequest) error {

	tx := dal.Gorm.Begin()

	// 删除之前所绑定的角色
	if err := tx.Model(model.UserRole{}).Where("user_id = ?", param.UserId).Delete(&model.UserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 绑定新角色
	for _, roleId := range param.RoleIds {
		if err := tx.Model(model.UserRole{}).Create(&model.UserRole{UserId: param.UserId, RoleId: roleId}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// 根据用户id集合查询角色id集合
func (u *UserRoleService) GetRoleIdsByUserIds(userIds []int) []int {

	var roleIds []int

	dal.Gorm.Model(model.UserRole{}).Where("user_id IN ?", userIds).Pluck("role_id", &roleIds)

	return roleIds
}

// 根据角色id集合查询用户id集合
func (u *UserRoleService) GetUserIdsByRoleIds(roleIds []int) []int {

	var userIds []int

	dal.Gorm.Model(model.UserRole{}).Where("role_id IN ?", roleIds).Pluck("user_id", &userIds)

	return userIds
}
