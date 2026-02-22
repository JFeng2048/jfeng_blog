package middleware

import (
	"jfeng_blog/internal/utils/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterCORS(r *gin.Engine) {
	corsConfig := config.GetCORSConfig()
	r.Use(cors.New(
		cors.Config{
			AllowOrigins:     corsConfig.AllowOrigins,
			AllowMethods:     corsConfig.AllowMethods,
			AllowHeaders:     corsConfig.AllowHeaders,
			ExposeHeaders:    corsConfig.ExposeHeaders,
			AllowCredentials: corsConfig.AllowCredentials,
			MaxAge:           12 * time.Hour,
		},
	))
}
