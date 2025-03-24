package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/robertantonyjaikumar/hangover-common/config"
)

var jwtSecret = config.CFG.V.GetString("jwt.secret")

// GenerateToken creates a new JWT token with expiration
func GenerateToken(userID string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// ParseToken verifies a token and extracts claims
func ParseToken(tokenString string) (*jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return &claims, nil
}
