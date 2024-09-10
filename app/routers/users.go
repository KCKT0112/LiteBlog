package routes

import (
	"github.com/KCKT0112/LiteBlog/app/controllers"
	"github.com/KCKT0112/LiteBlog/app/services"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	usersService := services.NewUsersService()
	usersController := controllers.NewUsersController(usersService)

	usersGroup := router.Group("/user")
	{
		usersGroup.GET("/profile", usersController.GetProfile)
	}
}
