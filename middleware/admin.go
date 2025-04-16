package middleware

import (
	error2 "Learning/error"
	"Learning/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": error2.TokenNotFound})
			return
		}

		user := helper.FetchUserFromToken(c.GetHeader("Authorization"))
		if user == nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch user data"})
			return
		}

		if !user.IsAdmin {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied. Admins only."})
			c.Abort()
			return
		}

		c.Next()
	}
}
