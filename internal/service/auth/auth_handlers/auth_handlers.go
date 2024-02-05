package auth_handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Zasedatelev/gin-rest-api.git/config"
	"github.com/Zasedatelev/gin-rest-api.git/internal/models"
	"github.com/Zasedatelev/gin-rest-api.git/internal/repository"
	"github.com/Zasedatelev/gin-rest-api.git/internal/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AuthHandler struct {
	DB *repository.Repository
}

func NewAuthHandler(DB *repository.Repository) AuthHandler {
	return AuthHandler{DB}
}

// RegistrUser godoc
// @Summary      register user
// @Description  post user
// @Tags         auth
// @Param        payload body models.SingUpInput true  "form data"
// @Success 200 {string} string "OK"
// @Success 400 {string} string "BAD REQUEST"
// @Success 404 {string} string "NOT FOUND"
// @Router       /auth/register [post]
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

	_, err = stmt.Exec(payload.Name, strings.ToLower(payload.Email), hashPass)
	if err != nil {
		log.WithError(err).Error("ОШИБКА ПРИ РЕГИСТРАЦИИ")
		c.JSON(http.StatusConflict, gin.H{"message": err})
	} else {
		log.WithFields(log.Fields{
			"user":    payload.Name,
			"created": time.Now(),
		}).Info("НОВЫЙ ПОЛЬЗОВАТЕЛЬ ДОБАВЛЕН")
		c.JSON(http.StatusCreated, gin.H{"massege": "New user account registered"})
		return
	}

}

// Login godoc
// @Summary      login user
// @Description  login user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload body models.SingInInput true  "form data"
// @Success 200 {string} string "OK"
// @Success 400 {string} string "BAD REQUEST"
// @Success 404 {string} string "NOT FOUND"
// @Router       /api/auth/login [post]
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid name or password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Password"})
		return
	}

	token, err := utils.GenerateToken(config.JWT.AccessTokenExpireDuration, user.ID, config.JWT.Secret)
	if err != nil {
		log.WithError(err).Error("ОШИБКА ПРИ ГЕНЕРАЦИИ ТОКЕНА")
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.SetCookie("token", token, 60, "/", "localhost", false, true)
	log.WithFields(log.Fields{
		"in": time.Now(),
	}).Info("ВХОД ВЫПОЛНЕН")
	c.JSON(http.StatusOK, gin.H{"status": "success", "token": token})
}
