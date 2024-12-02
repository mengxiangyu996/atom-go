package service

import (
	"atom-go/internal/dto"
	"atom-go/internal/model"
	"atom-go/pkg/dal"
)

// 菜单权限服务
type MenuService struct{}

// 创建菜单权限
func (m *MenuService) CreateMenu(menu *dto.SaveMenuRequest) error {

	data := model.Menu{
		MenuName:  menu.MenuName,
		ParentId:  menu.ParentId,
		MenuType:  menu.MenuType,
		Sort:      menu.Sort,
		Component: menu.Component,
		Icon:      menu.Icon,
		Path:      menu.Path,
		Method:    menu.Method,
		Status:    menu.Status,
		CreateBy:  menu.CreateBy,
	}

	return dal.Gorm.Model(model.Menu{}).Create(&data).Error
}

// 根据菜单权限id删除菜单权限
func (m *MenuService) DeleteMenuByMenuId(menuId int) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.Menu{}).Where("menu_id = ?", menuId).Delete(&model.Menu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除角色绑定的菜单权限
	if err := tx.Model(model.RoleMenu{}).Where("menu_id = ?", menuId).Delete(&model.RoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新菜单权限
func (m *MenuService) UpdateMenuByMenuId(menu *dto.SaveMenuRequest) error {

	data := model.Menu{
		MenuName:  menu.MenuName,
		ParentId:  menu.ParentId,
		MenuType:  menu.MenuType,
		Sort:      menu.Sort,
		Component: menu.Component,
		Icon:      menu.Icon,
		Path:      menu.Path,
		Method:    menu.Method,
		Status:    menu.Status,
		UpdateBy:  menu.UpdateBy,
	}

	return dal.Gorm.Model(model.Menu{}).Where("menu_id = ?", menu.MenuId).Updates(&data).Error
}

// 根据菜单资源id获取菜单资源信息
func (m *MenuService) GetMenuInfoByMenuId(menuId int) dto.MenuInfoResponse {

	var menu dto.MenuInfoResponse

	dal.Gorm.Model(model.Menu{}).Where("menu_id = ?", menuId).Take(&menu)

	return menu
}

// 根据路径和方法获取菜单权限
func (m *MenuService) GetMenuByPathAndMethod(path string, method string) dto.MenuInfoResponse {

	var menu dto.MenuInfoResponse

	dal.Gorm.Model(model.Menu{}).Where("path = ? AND method = ?", path, method).Last(&menu)

	return menu
}

// 获取菜单权限列表
// 菜单类型：M目录 C菜单 B按钮
func (m *MenuService) GetMenuList(param *dto.GetMenuListRequest) []dto.MenuListResponse {

	menus := make([]dto.MenuListResponse, 0)

	tx := dal.Gorm.Model(model.Menu{}).Order("sort DESC, menu_id")

	if len(param.MenuIds) > 0 {
		tx.Where("menu_id IN ?", param.MenuIds)
	}

	if param.Status != "" {
		tx.Where("status = ?", param.Status)
	}

	if len(param.MenuType) > 0 {
		tx.Where("menu_type IN ?", param.MenuType)
	}

	if param.MenuName != "" {
		tx.Where("menu_name LIKE ?", "%"+param.MenuName+"%")
	}

	tx.Find(&menus)

	return menus
}

// 菜单权限列表转菜单列表树形
func (m *MenuService) MenuListToTree(menus []dto.MenuListResponse, parentId int) []dto.MenuTreeResponse {

	tree := make([]dto.MenuTreeResponse, 0)

	for _, menu := range menus {
		if menu.ParentId == parentId {
			tree = append(tree, dto.MenuTreeResponse{
				MenuListResponse: menu,
				Children:         m.MenuListToTree(menus, menu.MenuId),
			})
		}
	}

	return tree
}
