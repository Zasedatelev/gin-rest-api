package auth_handlers

import (
	"net/http"
	"strings"

	"github.com/Oleg-OMON/gin-rest-api.git/internal/repository"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/service/auth/auth_models"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	DB *repository.Repository
}

func NewAuthHandler(DB *repository.Repository) AuthHandler {
	return AuthHandler{DB}
}
func (a *AuthHandler) RegistrUser(c *gin.Context) {
	payload := auth_models.SingUpInput{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	addQuery := `INSERT INTRO users (name, email, password) VALUES (1$, 2$, 3$)`
	_, err := a.DB.DataBase.Exec(addQuery, payload.Name, strings.ToLower(payload.Email), payload.Password)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "User added"})
}
