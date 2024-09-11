package middleware

import (
	"github.com/gin-gonic/gin"
)

func DefaultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Request", "Middleware")
		c.Next()
	}
}
