package api_v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterV1(r *gin.RouterGroup) {
	// 用户路由
	RegisterUserRoutes(r)
}
