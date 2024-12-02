package validator

import (
	"atom-go/internal/common/utils"
	"atom-go/internal/constant/regexp"
	"atom-go/internal/dto"
	"errors"
)

// 创建用户验证器
func CreateUserValidator(param *dto.SaveUserRequest) error {

	if param.UserType == "" {
		return errors.New("请选择用户类型")
	}

	if param.Username == "" {
		return errors.New("请输入用户账号")
	}

	if param.Password == "" {
		return errors.New("请输入用户密码")
	}

	if param.Nickname == "" {
		return errors.New("请输入用户昵称")
	}

	if param.Email != "" && !utils.CheckRegex(regexp.Email, param.Email) {
		return errors.New("请输入正确的邮箱")
	}

	if param.Phone != "" && !utils.CheckRegex(regexp.Phone, param.Phone) {
		return errors.New("请输入正确的手机号")
	}

	return nil
}

// 更新用户验证器
func UpdateUserValidator(param *dto.SaveUserRequest) error {

	if param.UserId <= 0 {
		return errors.New("参数错误")
	}

	if param.Email != "" && !utils.CheckRegex(regexp.Email, param.Email) {
		return errors.New("请输入正确的邮箱")
	}

	if param.Phone != "" && !utils.CheckRegex(regexp.Phone, param.Phone) {
		return errors.New("请输入正确的手机号")
	}

	return nil
}
