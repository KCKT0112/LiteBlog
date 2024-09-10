package utils

import (
	"time"

	"github.com/KCKT0112/LiteBlog/app/config"
	"github.com/golang-jwt/jwt/v5"
)

// Custom Claims
type Claims struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func GenerateJWT(id string) (string, error) {
	jwtKey := []byte(config.AppConfig.Auth.JwtSecret)
	// Set expiration time
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	jwtKey := []byte(config.AppConfig.Auth.JwtSecret)
	claims := &Claims{}

	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// 校验token是否合法
	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
