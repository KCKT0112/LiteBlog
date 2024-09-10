package controllers

import (
	"net/http"

	"github.com/KCKT0112/LiteBlog/app/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService services.UsersService
}

func NewUsersController(service services.UsersService) *UsersController {
	return &UsersController{
		usersService: service,
	}
}

func (uc *UsersController) GetProfile(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("it works!"))
}
