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

// @Summary Register
// @Description Register
// @Tags    Auth
// @Accept  json
// @Produce  json
// @Param   form  body  models.UserRegisterForm  true  "Register form"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /auth/register [post]
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

// @Summary Login
// @Description Login
// @Tags    Auth
// @Accept  json
// @Produce  json
// @Param   form  body  models.UserLoginForm  true  "Login form"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /auth/login [post]
func (u *UserController) Login(c *gin.Context) {
	var form models.UserLoginForm
	// bind form to struct
	if err := c.ShouldBind(&form); err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return

	}

	// check if user exists
	user, err := u.userService.GetUserByEmail(form.Email)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	if user == nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: "user not found"})
		return
	}

	// check if password is correct
	if user.Password != fmt.Sprintf("%x", sha256.Sum256([]byte(form.Password+config.AppConfig.Auth.PasswordSalt))) {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: "password is incorrect"})
		return
	}

	// generate token
	access_token, err := utils.GenerateAccessToken(user.ID.String())
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	refresh_token, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	// generate response data
	res_data := map[string]interface{}{
		"access_token":  access_token,
		"refresh_token": refresh_token,
		"user": map[string]interface{}{
			"id":   user.ID,
			"name": user.Name,
		},
	}

	// return token
	c.JSON(http.StatusOK, utils.Response{Code: 200, Message: "success", Data: res_data})
}

// @Summary Refresh Token
// @Description Refresh Token
// @Tags    Auth
// @Accept  json
// @Produce  json
// @Param   form  body  models.UserRefreshTokenForm  true  "Refresh Token form"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /auth/refresh [post]
func (u *UserController) RefreshToken(c *gin.Context) {
	var form models.UserRefreshTokenForm
	// bind form to struct
	if err := c.ShouldBind(&form); err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	// check if token is valid
	claims, err := utils.ValidateJWT(form.RefreshToken)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	// generate new token
	access_token, err := utils.GenerateAccessToken(claims.Id)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	refresh_token, err := utils.GenerateRefreshToken(claims.Id)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Response{Code: 400, Message: err.Error()})
		return
	}

	// generate response data
	res_data := map[string]interface{}{
		"access_token":  access_token,
		"refresh_token": refresh_token,
	}

	// return token
	c.JSON(http.StatusOK, utils.Response{Code: 200, Message: "success", Data: res_data})
}
