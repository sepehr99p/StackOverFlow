package handlers

import (
	"Learning/models"
	"Learning/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// LoginHandler
// @Tags auth
// @Accept json
// @Produce json
// @Param answer body models.User true "User object"
// @Success 201 {object} models.User
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}
	fmt.Printf("The user request value %v", user)

	if user.UserName == "Check" && user.Password == "123456" {
		tokenString, err := token.CreateToken(user.UserName)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "no username found"})
			return
		}
		c.IndentedJSON(http.StatusOK, tokenString)
		return
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	}
}

// ProtectedHandler
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} string
// @Router /protected [get]
func ProtectedHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization header"})
		return
	}

	const prefix = "Bearer "
	if !strings.HasPrefix(tokenString, prefix) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header format"})
		return
	}

	tokenString = strings.TrimPrefix(tokenString, prefix)

	err := token.VerifyToken(tokenString)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token", "error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "successful login"})
}
