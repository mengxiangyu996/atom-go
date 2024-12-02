package service

import (
	"atom-go/internal/dto"
	"atom-go/internal/model"
	"atom-go/pkg/dal"
)

// 用户服务
type UserService struct{}

// 创建用户
func (u *UserService) CreateUser(user *dto.SaveUserRequest) error {

	data := model.User{
		UserType: user.UserType,
		Username: user.Username,
		Password: user.Password,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Email:    user.Email,
		Phone:    user.Phone,
		Gender:   user.Gender,
		Status:   user.Status,
		CreateBy: user.CreateBy,
	}

	err := dal.Gorm.Model(model.User{}).Create(&data).Error

	user.UserId = data.UserId

	return err
}

// 根据用户id删除用户
func (u *UserService) DeleteUserByUserId(userId int) error {

	tx := dal.Gorm.Begin()

	if err := tx.Model(model.User{}).Where("user_id = ?", userId).Delete(&model.User{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除用户绑定的角色
	if err := tx.Model(model.UserRole{}).Where("user_id = ?", userId).Delete(&model.UserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// 更新用户
func (u *UserService) UpdateUserByUserId(user *dto.SaveUserRequest) error {

	data := model.User{
		UserType:      user.UserType,
		Username:      user.Username,
		Password:      user.Password,
		Nickname:      user.Nickname,
		Avatar:        user.Avatar,
		Email:         user.Email,
		Phone:         user.Phone,
		Gender:        user.Gender,
		LoginLastIp:   user.LoginLastIp,
		LoginLastTime: user.LoginLastTime,
		Status:        user.Status,
		CreateBy:      user.CreateBy,
		UpdateBy:      user.UpdateBy,
	}

	return dal.Gorm.Model(model.User{}).Where("user_id = ?", user.UserId).Updates(&data).Error
}

// 根据用户id获取用户信息
func (u *UserService) GetUserInfoByUserId(userId int) dto.UserInfoResponse {

	var user dto.UserInfoResponse

	dal.Gorm.Model(model.User{}).Where("user_id = ?", userId).Take(&user)

	return user
}

// 根据用户名获取用户信息
func (u *UserService) GetUserInfoByUsername(username string) dto.UserTokenResponse {

	var user dto.UserTokenResponse

	dal.Gorm.Model(model.User{}).Where("username = ?", username).Take(&user)

	return user
}

// 用户列表
func (u *UserService) GetUserPage(param *dto.GetUserPageRequest) ([]dto.UserListResponse, int) {

	var count int64
	users := make([]dto.UserListResponse, 0)

	tx := dal.Gorm.Model(model.User{}).Order("user_id DESC")

	if param.UserType != "" {
		tx.Where("user_type = ?", param.UserType)
	}

	if param.Username != "" {
		tx.Where("username = ?", param.Username)
	}

	if param.Nickname != "" {
		tx.Where("nickname = ?", param.Nickname)
	}

	if param.Status != "" {
		tx.Where("status = ?", param.Status)
	}

	tx.Count(&count).Offset((param.Page - 1) * param.PageSize).Limit(param.PageSize).Find(&users)

	return users, int(count)
}
