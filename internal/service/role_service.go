package service

import (
	"atom-go/internal/dto"
	"atom-go/internal/model"
	"atom-go/pkg/dal"
)

// 角色服务
type RoleService struct{}

// 创建角色
func (r *RoleService) CreateRole(role *dto.SaveRoleRequest) error {

	data := model.Role{
		RoleName: role.RoleName,
		Status:   role.Status,
		CreateBy: role.CreateBy,
	}

	return dal.Gorm.Model(model.Role{}).Create(&data).Error
}

// 根据角色id删除角色
func (r *RoleService) DeleteRoleByRoleId(roleId int) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.Role{}).Where("role_id = ?", roleId).Delete(&model.Role{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除角色绑定的菜单权限
	if err := tx.Model(model.RoleMenu{}).Where("role_id = ?", roleId).Delete(&model.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新角色
func (r *RoleService) UpdateRoleByRoleId(role *dto.SaveRoleRequest) error {

	data := model.Role{
		RoleName: role.RoleName,
		Status:   role.Status,
		UpdateBy: role.UpdateBy,
	}

	return dal.Gorm.Model(model.Role{}).Where("role_id = ?", role.RoleId).Updates(&data).Error
}

// 根据角色id获取角色信息
func (r *RoleService) GetRoleInfoByRoleId(roleId int) dto.RoleInfoResponse {

	var role dto.RoleInfoResponse

	dal.Gorm.Model(model.Role{}).Where("role_id = ?", roleId).Take(&role)

	return role
}

// 根据角色名称获取角色信息
func (r *RoleService) GetRoleInfoByRoleName(roleName string) dto.RoleInfoResponse {

	var role dto.RoleInfoResponse

	dal.Gorm.Model(model.Role{}).Where("role_name = ?", roleName).Take(&role)

	return role
}

// 角色列表
func (r *RoleService) GetRoleList(param *dto.GetRolePageRequest) ([]dto.RoleListResponse, int) {

	var count int64
	roles := []dto.RoleListResponse{}

	tx := dal.Gorm.Model(model.Role{}).Order("role_id DESC")

	if param.RoleName != "" {
		tx.Where("role_name LIKE ?", "%"+param.RoleName+"%")
	}

	if param.Status != "" {
		tx.Where("status = ?", param.Status)
	}

	tx.Count(&count).Offset((param.Page - 1) * param.PageSize).Limit(param.PageSize).Find(&roles)

	return roles, int(count)
}
