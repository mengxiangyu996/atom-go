package validator

import (
	"atom-go/internal/dto"
	"errors"
)

// 创建角色验证器
func CreateRoleValidator(param *dto.SaveRoleRequest) error {

	if param.RoleName == "" {
		return errors.New("请输入角色名称")
	}

	return nil
}

// 更新角色验证器
func UpdateRoleValidator(param *dto.SaveRoleRequest) error {

	if param.RoleId <= 0 {
		return errors.New("参数错误")
	}

	return nil
}
