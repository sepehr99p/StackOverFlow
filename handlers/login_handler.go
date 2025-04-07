package handlers

import (
	"Learning/database"
	"Learning/models"
	"Learning/token"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

func isUserAlreadyExist(phoneNumber string) bool {
	var count int64
	database.DB.Model(&models.User{}).
		Where("user_name = ?", phoneNumber).
		Count(&count)
	return count > 0
}

func createUser(user models.UserRegister) *string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		message := "Error hashing password"
		return &message
	}
	newUser := models.User{
		UserName: user.PhoneNumber,
		Password: string(hashedPassword),
	}
	result := database.DB.Create(&newUser)
	if result.Error != nil {
		message := result.Error.Error()
		return &message
	}
	return nil
}

// RegisterHandler
// @Tags auth
// @Accept json
// @Produce json
// @Param answer body models.UserRegister true "User object"
// @Success 201 {object} string
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var userInput models.UserRegister
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON format", "error": err.Error(),
		})
		return
	}

	if isUserAlreadyExist(userInput.PhoneNumber) {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	}

	var userCreatingResult = createUser(userInput)
	if userCreatingResult != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": userCreatingResult})
	}

	tokenString, err := token.CreateToken(userInput.PhoneNumber)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	c.IndentedJSON(http.StatusCreated, tokenString)
}

// LoginHandler
// @Tags auth
// @Accept json
// @Produce json
// @Param answer body models.UserRegister true "User object"
// @Success 201 {object} string
// @Router /login [post]
func LoginHandler(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	var user models.UserRegister
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON format"})
		return
	}

	dbUser := models.User{UserName: user.PhoneNumber, Password: user.Password}
	if queryResult := database.DB.Find(&dbUser).Error; queryResult != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dbUser.Password))
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	tokenString, err := token.CreateToken(user.PhoneNumber)
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
