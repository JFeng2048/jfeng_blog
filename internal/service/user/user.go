package user

import (
	"errors"

	"go.uber.org/zap"

	"jfeng_blog/internal/repository"
	"jfeng_blog/internal/model"
)

func GetUserInfoByID(id int) (*model.User, error) {
	// 处理获取用户信息逻辑
	if id <= 0 {
		zap.L().Error("无效的用户ID", zap.Int("id", id))
		return nil, errors.New("无效的用户ID")
	}

	user, err := repository.GetUserById(uint(id))
	if err != nil {
		zap.L().Error("获取用户信息失败", zap.Int("id", id), zap.Error(err))
		return nil, err
	}
	zap.L().Info("成功获取用户信息", zap.Int("id", id))
	return user, nil
}