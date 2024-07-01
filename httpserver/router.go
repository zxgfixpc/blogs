package httpserver

import (
	"blogs/controller"
	"blogs/middleware"

	"github.com/gin-gonic/gin"
)

func registerRouter(router *gin.Engine) {
	notLoginGroup := router.Group("blogs/")
	{
		notLoginGroup.POST("login", controller.Login)
		notLoginGroup.GET("recommend-article", controller.GetRecommendArticles)
	}

	needLoginGroup := router.Group("blogs/").Use(middleware.LoginMiddleware())
	{
		needLoginGroup.POST("exit", controller.Exit)
		needLoginGroup.POST("article-create-or-update", controller.CreateOrUpdateArticle)
	}

}
