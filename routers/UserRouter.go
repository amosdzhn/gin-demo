package routers

import (
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(engine *gin.Engine) {
	userRouter(engine)
}

func userRouter(engine *gin.Engine) {
	userRouterGroup := engine.Group("/user")
	{
		userController := &controller.UserController{}
		userRouterGroup.GET("/hello", userController.Hello)
	}

	configRouterGroup := engine.Group("/config")
	{
		configController := &controller.ConfigController{}
		configRouterGroup.GET("/info", configController.GetConfig)
	}
}
