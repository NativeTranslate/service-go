package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SecretKey = []byte("secret")

// GenerateToken generates a new token with the given username
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// ParseToken parses a token string and returns the username
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	return username, nil
}
