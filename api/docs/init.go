package docs

import (
	"strconv"

	"go.uber.org/zap"
	

	"jfeng_blog/internal/utils/config"
)

func InitSwagger() {
   
	zap.L().Info("初始化 Swagger 文档")
    conf := config.GetConfig()
	SwaggerInfo.Host = conf.Server.Host + ":" + strconv.Itoa(conf.Server.Port)
    SwaggerInfo.Version = conf.System.Version
	SwaggerInfo.Title = conf.Server.Name
	zap.L().Info("Swagger 文档初始化完成", zap.String("host", SwaggerInfo.Host), zap.String("version", SwaggerInfo.Version), zap.String("title", SwaggerInfo.Title))
}
