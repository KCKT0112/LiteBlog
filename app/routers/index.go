package routes

import (
	"github.com/KCKT0112/LiteBlog/app/controllers"
	"github.com/KCKT0112/LiteBlog/app/services"
	"github.com/gin-gonic/gin"
)

func IndexRoutes(router *gin.RouterGroup) {
	indexService := services.NewIndexService()
	indexController := controllers.NewIndexController(indexService)

	indexGroup := router.Group("/")
	{
		indexGroup.GET("/", indexController.GetIndex)
	}
}
