package handlers

import (
	"Learning/database"
	"Learning/models"
	"Learning/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// RegisterHandler
// @Tags auth
// @Accept json
// @Produce json
// @Param answer body models.UserRegister true "User object"
// @Success 201 {object} string
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var user models.UserRegister
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format", "error": err.Error()})
		return
	}
	var userModel = models.User{UserName: user.PhoneNumber, Password: user.Password}
	dataExistenceResult := database.DB.Find(&userModel)
	if dataExistenceResult.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error querying database"})
		return
	}
	if dataExistenceResult.RowsAffected == 0 {
		result := database.DB.Create(&user)
		if result.Error != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating user", "error": result.Error.Error()})
			return
		}
		tokenString, err := token.CreateToken(userModel.UserName)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			return
		}
		c.IndentedJSON(http.StatusCreated, tokenString)
	} else {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "user already exist"})
		return
	}

}

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
	dbUser := models.User{UserName: user.UserName, Password: user.Password}
	if queryResult := database.DB.Find(&dbUser).Error; queryResult != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	tokenString, err := token.CreateToken(user.UserName)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	c.IndentedJSON(http.StatusOK, tokenString)
	return
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
