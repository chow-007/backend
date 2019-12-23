package main

import (
	"backend/configs"
	"backend/router"
	"backend/services"
	"github.com/gin-gonic/gin"
)


// @title Docker监控服务
// @version 1.0
// @description docker监控服务后端API接口文档

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host 127.0.0.1:9001
// @BasePath
func main() {
	defer func() {
		services.Service.Close()
	}()
	if configs.Default.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := router.InitRouter()
	router.Run(":9001")
}
