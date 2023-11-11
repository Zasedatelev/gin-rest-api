package main

import (
	"github.com/Oleg-OMON/gin-rest-api.git/cmd/app/docs"
	_ "github.com/Oleg-OMON/gin-rest-api.git/cmd/app/docs"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/repository"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/routers"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/service/auth/auth_handlers"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/service/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var server *gin.Engine

func init() {

}

// @title           Test Golang API
// @version         1.20
// @description     Привет. Мой не большой проект для изучения програмирования API на Go.
// @contact.name   Oleg Zasedatelev
// @host      localhost:8080
// @BasePath  /api
func main() {
	repo := new(repository.Repository)
	repo.Open()

	GameHandler := handlers.NewGameHandler(repo)
	GameRouter := routers.NewGameRouteController(GameHandler)

	AuthHandler := auth_handlers.NewAuthHandler(repo)
	AuthRoutrer := routers.NewAuthRouteController(AuthHandler)

	server = gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router := server.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"

	GameRouter.InitGameRouters(router)
	AuthRoutrer.InitAuthRouters(router)
	server.Run()

}
