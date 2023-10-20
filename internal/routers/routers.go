package routers

import (
	"github.com/Oleg-OMON/gin-rest-api.git/internal/handlers"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/store"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	router *gin.Engine
}

func InitRouters(s *store.Store) {
	handler := handlers.Handler{
		DB: s,
	}
	router := gin.Default()

	v1 := router.Group("v1/")
	{
		v1.GET("/all_players", handler.AllPlayers)
		v1.GET("/all_games", handler.AllGames)
		v1.GET("/results_games/:nickname", handler.ResultGames)

	}

	router.Run()
}
