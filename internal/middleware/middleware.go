package middleware

import (
	"net/http"
	"strings"

	"github.com/Oleg-OMON/gin-rest-api.git/config"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/utils"
	"github.com/gin-gonic/gin"
)

// Проверка подлиности токена
func AuthorizationUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		config := config.LoadConfig()
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

		sub, err := utils.ValidateToken(token, config.JWT.Secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		c.Set("UserID", sub)
		c.Next()

	}
}
