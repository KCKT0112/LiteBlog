package controllers

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/KCKT0112/LiteBlog/app/config"
	"github.com/KCKT0112/LiteBlog/app/models"
	"github.com/KCKT0112/LiteBlog/app/services"
	"github.com/KCKT0112/LiteBlog/app/utils"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		userService: service,
	}
}

func (u *UserController) Profile(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("it works!"))
}

func (u *UserController) Register(c *gin.Context) {
	var form models.UserRegisterForm
	// bind form to struct
	if err := c.ShouldBind(&form); err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	// create user
	user := models.User{
		Name:      form.Username,
		Email:     form.Email,
		Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(form.Password+config.AppConfig.Auth.PasswordSalt))), // Hash password
		CreatedAt: int64(time.Now().Unix()),                                                                   // time stamp
		UpdatedAt: int64(time.Now().Unix()),
	}

	// check if user exists
	if _u, _ := u.userService.GetUserByEmail(user.Email); _u != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: "user already exists"})
		return
	}

	result, err := u.userService.CreateUser(user)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.Response{Code: 200, Message: "success", Data: result})
}

func (u *UserController) Login(c *gin.Context) {

}
