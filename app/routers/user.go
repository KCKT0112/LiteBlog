package routes

import (
	"github.com/KCKT0112/LiteBlog/app/controllers"
	"github.com/KCKT0112/LiteBlog/app/services"
	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	usersGroup := router.Group("/user")
	{
		usersGroup.GET("/profile", userController.Profile)
	}
}
