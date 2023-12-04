package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

// Создание токена
func GenerateToken(ttl time.Duration, payload interface{}, secretKey string) (string, error) {

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub": payload,
		"exp": now.Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Error("JWT token signing")
	}

	return t, nil
}

func ValidateToken(token string, signedJWTKey string) (interface{}, error) {
	t, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["arg"])
		}

		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil

}
