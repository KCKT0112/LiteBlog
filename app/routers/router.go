package routes

import (
	"github.com/KCKT0112/LiteBlog/app/middleware"
	docs "github.com/KCKT0112/LiteBlog/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	// Register middleware
	router.Use(middleware.DefaultMiddleware())

	// Register routes
	api := router.Group("/api")
	{
		IndexRoutes(api)
		AuthRoutes(api)
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			UsersRoutes(protected)
		}
	}

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
