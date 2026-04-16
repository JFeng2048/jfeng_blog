package repository

import (
	"jfeng_blog/internal/model"
	"jfeng_blog/internal/utils/config"
)

func GetPostInfoByID(postID string) (*model.Post, error) {
	var post model.Post
	if err := config.GetDB().Where("id = ?", postID).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}