package helper

import (
	"Learning/database"
	"Learning/models"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func FetchUserFromToken(header string) *models.User {
	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		return nil
	}
	tokenString := strings.TrimPrefix(header, "Bearer ")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	if err != nil {
		return nil
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil
	}
	var user models.User
	if err := database.DB.Where("user_name = ?", username).First(&user).Error; err != nil {
		return nil
	}
	return &user
}
