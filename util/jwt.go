package util

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(username string) (string, error) {
	secret := []byte(os.Getenv("SECRET"))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                         // Subject (user identifier)
		"iss": "todo-app",                       // Issuer
		"aud": username,                         // Audience (user role)
		"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                // Issued at
	})

	token, err := claims.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	secret := []byte(os.Getenv("SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Return the verified token
	return token, nil
}
