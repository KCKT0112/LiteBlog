package middleware

import (
	"net/http"
	"strings"

	"github.com/KCKT0112/LiteBlog/app/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, utils.Error(401, "authorization header missing"))
			c.Abort()
			return
		}

		// validate token
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.Error(401, "invalid token"))
			c.Abort()
			return
		}

		// set userID in context
		c.Set("userID", claims.ID)

		c.Next()
	}
}
