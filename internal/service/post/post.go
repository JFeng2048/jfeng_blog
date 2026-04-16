package post

import (
	"jfeng_blog/internal/model"
	"jfeng_blog/internal/repository"
)

func GetPostInfoByID(postID string) (*model.Post, error) {
	post, err := repository.GetPostInfoByID(postID)
	if err != nil {
		return nil, err
	}
	return post, nil
}