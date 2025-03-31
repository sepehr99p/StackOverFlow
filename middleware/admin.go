package middleware

import (
	"Learning/database"
	"Learning/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := strconv.Atoi(c.GetHeader("User-ID")) // user id should be in header for this to work

		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
			c.Abort()
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
