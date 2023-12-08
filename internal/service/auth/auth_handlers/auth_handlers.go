package auth_handlers

import (
	"net/http"
	"strings"

	"github.com/Oleg-OMON/gin-rest-api.git/config"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/models"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/repository"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AuthHandler struct {
	DB *repository.Repository
}

func NewAuthHandler(DB *repository.Repository) AuthHandler {
	return AuthHandler{DB}
}
func (a *AuthHandler) RegistrUser(c *gin.Context) {
	payload := models.SingUpInput{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	addQuery := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	stmt, err := a.DB.DataBase.Prepare(addQuery)
	if err != nil {
		log.Debug(err)
	}
	defer stmt.Close()

	hashPass, err := utils.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	result, err := stmt.Exec(payload.Name, strings.ToLower(payload.Email), hashPass)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	rowId, err := result.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{"lastId": rowId, "massege": "New user account registered"})
}

func (a *AuthHandler) Login(c *gin.Context) {
	payload := models.SingInInput{}
	config := config.LoadConfig()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := models.User{}

	err := a.DB.DataBase.Get(&user, `SELECT * FROM users WHERE name = $1`, payload.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Password"})
		return
	}

	token, err := utils.GenerateToken(config.JWT.AccessTokenExpireDuration, user.ID, config.JWT.Secret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.SetCookie("token", token, 60, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}
