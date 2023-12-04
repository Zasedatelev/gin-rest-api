package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Oleg-OMON/gin-rest-api.git/internal/models"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/repository"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/utils"
	"github.com/gin-gonic/gin"
)

type ConnectDB struct {
	DB *repository.Repository
}

func NewDB(DB *repository.Repository) ConnectDB {
	return ConnectDB{DB}
}

// Проверка подлиности токена
func (a *ConnectDB) authorizationUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		cookie, err := c.Cookie("token")
		header := c.Request.Header.Get("Authorization")
		fields := strings.Fields(header)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err == nil {
			token = cookie
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "You are not logged in"})
			return
		}

		sub, err := utils.ValidateToken(token, "jwt-token-secret")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.User
		result := a.DB.DataBase.Get(&user, `SELECT * FROM users WHERE name = $1`, fmt.Sprint(sub))
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		c.Set("currentUser", user)
		c.Next()

	}
}
