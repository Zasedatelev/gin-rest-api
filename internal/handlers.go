package internal

import (
	"fmt"
	"net/http"

	"github.com/Oleg-OMON/gin-rest-api.git/models"
	"github.com/Oleg-OMON/gin-rest-api.git/store"
	"github.com/gin-gonic/gin"
)

func All_players(c *gin.Context) {
	var store *store.Store = &store.Store{}
	rows, err := store.DataBase.Query("Select * From players")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	players := []models.Player{}

	for rows.Next() {
		pl := models.Player{}
		err := rows.Scan(&pl.Player_id, &pl.First_name, &pl.Last_name, &pl.Nickname, &pl.Citizenship, &pl.Dob, &pl.Role)
		if err != nil {
			fmt.Println("хуня 7 а не 6")
			continue
		}
		players = append(players, pl)
	}

	for _, player := range players {
		c.JSON(http.StatusOK, player)
	}

}
