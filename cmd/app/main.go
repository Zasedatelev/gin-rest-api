package main

import (
	"github.com/Oleg-OMON/gin-rest-api.git/internal"
	"github.com/Oleg-OMON/gin-rest-api.git/store"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	s := new(store.Store)
	s.Open()
	handler := internal.Handler{
		DB: s,
	}
	router.GET("/all_players", handler.AllPlayers)
	router.GET("/all_games", handler.AllGames)

	router.Run()
}
