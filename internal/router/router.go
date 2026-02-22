package router

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"jfeng_blog/api/docs"
	_ "jfeng_blog/api/docs"

	"jfeng_blog/internal/router/api_v1"
	"jfeng_blog/internal/utils/config"
	"jfeng_blog/internal/middleware"
)

func initRouter() *gin.Engine {
	zap.L().Info("初始化路由系统")
	r := gin.New()
	middleware.RegisterMiddleware(r)
	// 注册路由
	registerRoutes(r)
	zap.L().Info("路由系统初始化完成")
	return r
}

func registerRoutes(r *gin.Engine) {
	// 注册路由
	version := config.GetSystemConfig().Version

	parts := strings.Split(version, ".")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	if len(parts) != 3 {
		zap.L().Error("版本格式错误")
		return
	}

	apiGroup := r.Group("/api")
	{
		switch parts[0] {
		case "1":
			v1Group := apiGroup.Group("/v1")
			api_v1.RegisterV1(v1Group)
		default:
			v1Group := apiGroup.Group("/v1")
			api_v1.RegisterV1(v1Group)
		}
	}

	zap.L().Info("注册 Swagger 路由")
	docs.InitSwagger()
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})
	r.GET("/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
	))
	zap.L().Info("Swagger 路由注册完成")

	RegisterStaticRoutes(r)
}

func RegisterStaticRoutes(r *gin.Engine) {
	zap.L().Info("注册静态资源路由")
	r.StaticFile("/swagger.json", "./api/docs/swagger.json")
	r.StaticFile("/swagger.yaml", "./api/docs/swagger.yaml")
	zap.L().Info("静态资源路由注册完成")
}

func Run() {
	r := initRouter()
	serverConfig := config.GetServerConfig()

	srv := &http.Server{
		Addr:    serverConfig.Host + ":" + strconv.Itoa(serverConfig.Port),
		Handler: r,
	}

	zap.L().Info("启动路由系统", zap.String("address", serverConfig.Host+":"+strconv.Itoa(serverConfig.Port)))
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("路由系统启动失败", zap.Error(err))
		}
	}()
	zap.L().Info("路由系统启动完成")

	zap.L().Info("服务启动完成")
	zap.L().Info("服务地址: " + "http://" + config.GetServerConfig().Host + ":" + strconv.Itoa(config.GetServerConfig().Port))
	zap.L().Info("swagger 地址: " + "http://" + config.GetServerConfig().Host + ":" + strconv.Itoa(config.GetServerConfig().Port) + "/docs")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("服务停止关闭中")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("服务停止关闭失败", zap.Error(err))
	}
	zap.L().Info("服务停止关闭完成")
}

func GetAPIBasePath(){

}