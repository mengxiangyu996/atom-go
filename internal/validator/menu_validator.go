package validator

import (
	"atom-go/internal/dto"
	"errors"
)

// 创建菜单权限验证器
func CreateMenuValidator(param *dto.SaveMenuRequest) error {

	if param.MenuName == "" {
		return errors.New("请输入菜单名称")
	}

	if param.MenuType == "" {
		return errors.New("请选择菜单类型")
	}

	return nil
}

// 更新菜单权限验证器
func UpdateMenuValidator(param *dto.SaveMenuRequest) error {

	if param.MenuId <= 0 {
		return errors.New("参数错误")
	}

	return nil
}
