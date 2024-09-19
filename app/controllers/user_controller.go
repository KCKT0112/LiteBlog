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

// @Summary		Profile
// @Description	Profile
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	utils.Response{code=int, message=string, data=models.ProfileResponse}
// @Router			/user/profile [get]
// @Security		Bearer
func (u *UserController) Profile(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID not found"})
		return
	}

	userID := userIDInterface.(string)
	user, err := u.userService.GetUserByUID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(500, err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, utils.Error(404, "user not found"))
		return
	}

	// user convert to models.UserProfile
	userProfile := models.ProfileResponse{
		Uid:   user.Uid,
		Name:  user.Name,
		Email: user.Email,
		Rules: user.Rules,
	}

	c.JSON(http.StatusOK, utils.Success(userProfile))
}

// @Summary		Register
// @Description	Register
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			form	body		models.UserRegisterForm	true	"Register form"
// @Success		200		{object}	utils.Response{code=int, message=string, data=nil}
// @Failure		400		{object}	utils.Response{code=int, message=string}
// @Router			/auth/register [post]
func (u *UserController) Register(c *gin.Context) {
	var form models.UserRegisterForm
	// bind form to struct
	if err := c.ShouldBind(&form); err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	// create user
	user := models.User{
		Uid:       utils.GenerateUUID(),
		Name:      form.Username,
		Email:     form.Email,
		Password:  fmt.Sprintf("%x", sha256.Sum256([]byte(form.Password+config.AppConfig.Auth.PasswordSalt))), // Hash password
		Rules:     []string{"user"},
		CreatedAt: int64(time.Now().Unix()), // time stamp
		UpdatedAt: int64(time.Now().Unix()),
	}

	// loop check uuid is unique
	for {
		if _u, _ := u.userService.GetUserByUID(user.Uid); _u == nil {
			break
		} else {
			user.Uid = utils.GenerateUUID()
		}
	}

	// check if user exists
	if _u, _ := u.userService.GetUserByEmail(user.Email); _u != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, "user already exists"))
		return
	}

	_, err := u.userService.CreateUser(user)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.Success(nil))
}

// @Summary		Login
// @Description	Login
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			form	body		models.UserLoginForm	true	"Login form"
// @Success		200		{object}	utils.Response{code=int, message=string, data=models.LoginResponse}
// @Failure		400		{object}	utils.Response{code=int, message=string}
// @Router			/auth/login [post]
func (u *UserController) Login(c *gin.Context) {
	var form models.UserLoginForm
	// bind form to struct
	if err := c.ShouldBind(&form); err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return

	}

	// check if user exists
	user, err := u.userService.GetUserByEmail(form.Email)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	if user == nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, "user not found"))
		return
	}

	// check if password is correct
	if user.Password != fmt.Sprintf("%x", sha256.Sum256([]byte(form.Password+config.AppConfig.Auth.PasswordSalt))) {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, "password is incorrect"))
		return
	}

	// generate token
	access_token, err := utils.GenerateAccessToken(user.Uid)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	refresh_token, err := utils.GenerateRefreshToken(user.Uid)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	// generate response data
	res_data := map[string]interface{}{
		"access_token":  access_token,
		"refresh_token": refresh_token,
		"user": map[string]interface{}{
			"uid":   user.Uid,
			"name":  user.Name,
			"rules": user.Rules,
		},
	}

	// return token
	c.JSON(http.StatusOK, utils.Success(res_data))
}

// @Summary		Refresh Token
// @Description	Refresh Token
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param			form	body		models.UserRefreshTokenForm	true	"Refresh Token form"
// @Success		200		{object}	utils.Response{code=int, message=string, data=models.RefreshTokenResponse}
// @Failure		400		{object}	utils.Response{code=int, message=string}
// @Router			/auth/refresh [post]
func (u *UserController) RefreshToken(c *gin.Context) {
	var form models.UserRefreshTokenForm
	// bind form to struct
	if err := c.ShouldBind(&form); err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	// check if token is valid
	claims, err := utils.ValidateJWT(form.RefreshToken)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	// generate new token
	access_token, err := utils.GenerateAccessToken(claims.ID)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	refresh_token, err := utils.GenerateRefreshToken(claims.ID)
	if err != nil {
		// return error
		c.JSON(http.StatusBadRequest, utils.Error(400, err.Error()))
		return
	}

	// generate response data
	res_data := map[string]interface{}{
		"access_token":  access_token,
		"refresh_token": refresh_token,
	}

	// return token
	c.JSON(http.StatusOK, utils.Success(res_data))
}
