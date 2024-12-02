package validator

import (
	"atom-go/internal/common/utils"
	"atom-go/internal/constant/regexp"
	"atom-go/internal/dto"
	"errors"
)

// 登录验证器
func LoginValidator(param *dto.LoginRequest) error {

	if param.Username == "" {
		return errors.New("请输入账号")
	}

	if param.Password == "" {
		return errors.New("请输入密码")
	}

	if param.CaptchaId == "" {
		return errors.New("验证码错误")
	}

	if param.CaptchaCode == "" {
		return errors.New("请输入验证码")
	}

	return nil
}

// 更新个人信息验证器
func UpdateInfoValidator(param *dto.SaveUserRequest) error {

	if param.Email != "" && !utils.CheckRegex(regexp.Email, param.Email) {
		return errors.New("请输入正确的邮箱")
	}

	if param.Phone != "" && !utils.CheckRegex(regexp.Phone, param.Phone) {
		return errors.New("请输入正确的手机号")
	}

	return nil
}
