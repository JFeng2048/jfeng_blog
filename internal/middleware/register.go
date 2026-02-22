package middleware

import (
	"jfeng_blog/internal/utils/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterMiddleware(r *gin.Engine) {
	// 注册中间件
	r.Use(logger.ZapLogger())
	r.Use(logger.ZapRecovery(true))
	zap.L().Info("注册中间件")
	RegisterCORS(r)
	zap.L().Info("注册中间件完成")
}