package httpserver

import (
	"blogs/controller"
	"blogs/middleware"

	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	loginGroup := router.Group("blogs/")
	{
		loginGroup.POST("login", controller.Login)
	}

	commGroup := router.Group("blogs/")
	commGroup.Use(middleware.LoginMiddleware())
	{

	}
}
