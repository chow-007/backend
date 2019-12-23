package router

import (
	"backend/controllers"
	"backend/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "successed",
		})
	})

	router.Use(middleware.Cors())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	loginApi := router.Group("auth/login")
	{
		loginApi.POST("", controllers.Login) // 用户登录
	}
	//router.Use(jwtAuthenticateMiddleware)

	userApi := router.Group("/user")
	{
		userApi.GET("/detail/:userId", controllers.UserDetail)       //用户详情
		userApi.GET("/list/:limit/:offset", controllers.GetUserList) //获取用户列表
		userApi.PUT("", controllers.UserUpdate)                      //更新用户
		userApi.POST("", controllers.UserCreate)                     //创建用户
		userApi.DELETE("/:userId", controllers.UserDelete)           //删除用户
	}

	return router
}

