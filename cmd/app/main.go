package main

import (
	"github.com/Oleg-OMON/gin-rest-api.git/internal/repository"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/routers"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/service/auth"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/service/handlers"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var server *gin.Engine

func init() {
	s := new(repository.Repository)
	s.Open()

	AuthHandler := auth.NewAuthHandler(s)
	routers.NewAuthRouteController(AuthHandler)

	GameHandler := handlers.NewGameHandler(s)
	routers.NewGameRouteController(GameHandler)

	server = gin.Default()
}

func main() {

	server.Run()

}
