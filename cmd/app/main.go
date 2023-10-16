package main

import (
	"github.com/Oleg-OMON/gin-rest-api.git/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/all_players", internal.All_players)

	router.Run()
}
