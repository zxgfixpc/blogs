package httpserver

import (
	"blogs/controller"

	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	blogsGroup := router.Group("blogs/")
	{
		blogsGroup.GET("hello", controller.Hello)
	}

}
