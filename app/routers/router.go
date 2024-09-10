package routes

import (
	"github.com/KCKT0112/LiteBlog/app"
	docs "github.com/KCKT0112/LiteBlog/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gin Swagger Example API
// @version         1.0
// @description     This is a sample server for a Swagger API with Gin.

// @host      localhost:8083
// @BasePath  /api
func InitRouter() *gin.Engine {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api"

	// Register middleware
	router.Use(app.MiddWare())

	// Register routes
	api := router.Group("/api")
	{
		IndexRoutes(api)
		AuthRoutes(api)
		UsersRoutes(api)
	}

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
