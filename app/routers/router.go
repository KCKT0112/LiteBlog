package routes

import (
	"github.com/KCKT0112/LiteBlog/app"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// Register middleware
	router.Use(app.MiddWare())

	// Register routes
	IndexRoutes(router)
	UsersRoutes(router)

	return router
}
