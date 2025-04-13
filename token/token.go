package token

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"Learning/database"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

// ExtractTokenFromHeader securely extracts the token from the Authorization header
func ExtractTokenFromHeader(authHeader string) (string, error) {
	const prefix = "Bearer "
	if len(authHeader) < len(prefix) {
		return "", fmt.Errorf("invalid authorization header")
	}

	if !strings.HasPrefix(authHeader, prefix) {
		return "", fmt.Errorf("invalid authorization header format")
	}
	token := authHeader[len(prefix):]

	if len(token) > 8192 { // Reasonable maximum token length
		return "", fmt.Errorf("token too long")
	}

	return token, nil
}

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	cacheKey := "token:" + username
	if err := database.CacheUserToken(ctx, cacheKey, tokenString); err != nil {
		return "", fmt.Errorf("failed to cache token: %v", err)
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	ctx := context.Background()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid token claims")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return fmt.Errorf("invalid username in token")
	}

	cacheKey := "token:" + username
	cachedToken, err := database.GetCachedUserToken(ctx, cacheKey)
	if err != nil {
		return fmt.Errorf("token not found in cache")
	}

	if cachedToken != tokenString {
		return fmt.Errorf("token mismatch")
	}

	return nil
}

func InvalidateToken(username string) error {
	ctx := context.Background()
	cacheKey := "token:" + username
	return database.DeleteCachedToken(ctx, cacheKey)
}
