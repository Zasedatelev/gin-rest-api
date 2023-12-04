package utils

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Errorf("could not hash password: %s", err)

	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashPass string, payloadPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(payloadPass))
}
