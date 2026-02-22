package repository


import (
	"jfeng_blog/internal/model"
	"jfeng_blog/internal/utils/config"
)

func GetUserById(id uint) (*model.User, error) {
	var user model.User
	if err := config.GetDB().Where("user_id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}