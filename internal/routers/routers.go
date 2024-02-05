package routers

import (
	"github.com/Zasedatelev/gin-rest-api.git/internal/handlers"
	"github.com/Zasedatelev/gin-rest-api.git/internal/middleware"
	"github.com/Zasedatelev/gin-rest-api.git/internal/service/auth/auth_handlers"
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

	v1.GET("/all_players", r.GameHandler.GetAllPlayers)
	v1.GET("/all_games", middleware.AuthorizationUser(), r.GameHandler.AllGames)
	v1.GET("/results_games/:nickname", r.GameHandler.ResultGames)
	v1.GET("/get_player/:nickname", r.GameHandler.GetPlayer)

}

// группа путей для регистрации и авторизации
func (a *AuthRouteController) InitAuthRouters(gr *gin.RouterGroup) {
	auth := gr.Group("/auth")
	{
		auth.POST("/register", a.AuthController.RegistrUser)
		auth.POST("/login", a.AuthController.Login)
	}
}
