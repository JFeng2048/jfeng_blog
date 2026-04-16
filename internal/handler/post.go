package handler

import (
	"github.com/gin-gonic/gin"
	"jfeng_blog/internal/model"
)

func GetPostInfoByID(c *gin.Context) {
	postID := c.Param("id")
	post := model.Post{
		ID: postID,
	}
	c.JSON(200, post)
}
