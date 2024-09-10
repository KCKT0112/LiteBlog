package routes

import (
	"github.com/KCKT0112/LiteBlog/app/controllers"
	"github.com/KCKT0112/LiteBlog/app/services"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	userService := services.NewUserService()
	userController := controllers.NewUserController(userService)

	usersGroup := router.Group("/auth")
	{
		usersGroup.POST("/login", userController.Login)
		usersGroup.POST("/register", userController.Register)
	}
}
