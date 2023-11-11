package routers

import (
	"github.com/Oleg-OMON/gin-rest-api.git/internal/service/auth/auth_handlers"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/service/handlers"
	"github.com/gin-gonic/gin"
)

type GameRouteController struct {
	GameHandler handlers.GameHandler
}

func NewGameRouteController(GameHandler handlers.GameHandler) GameRouteController {
	return GameRouteController{GameHandler}
}

type AuthRouteController struct {
	AuthController auth_handlers.AuthHandler
}

func NewAuthRouteController(AuthController auth_handlers.AuthHandler) AuthRouteController {
	return AuthRouteController{AuthController}
}

// группа путей для предсатвления данный о играх
func (r *GameRouteController) InitGameRouters(gr *gin.RouterGroup) {
	v1 := gr.Group("/games")

	v1.GET("/all_players", r.GameHandler.AllPlayers)
	v1.GET("/all_games", r.GameHandler.AllGames)
	v1.GET("/results_games/:nickname", r.GameHandler.ResultGames)

}

// группа путей для регистрации и авторизации
func (a *AuthRouteController) InitAuthRouters(gr *gin.RouterGroup) {
	auth := gr.Group("/auth")
	{
		auth.POST("/register", a.AuthController.RegistrUser)
	}
}
