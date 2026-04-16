package api_v1

import (
	"github.com/gin-gonic/gin"
	"jfeng_blog/internal/handler"
)

func RegisterPostRoutes(r *gin.RouterGroup) {
	postGroup := r.Group("/post")
	{
		postGroup.GET("/:id", handler.GetPostInfoByID)

	}
}	