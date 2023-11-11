package auth_models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json: "id"`
	Name     string    `json: "name"`
	Email    string    `json: "email"`
	Password string    `json: "password"`
}

type SingUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SingInInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password"  binding:"required"`
}
