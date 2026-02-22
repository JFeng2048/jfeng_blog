package api_v1

import (
	"github.com/gin-gonic/gin"
	"jfeng_blog/internal/handler"
)



func RegisterUserRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", handler.GetUserInfoByID)
	}
}