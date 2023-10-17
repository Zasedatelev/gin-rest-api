package internal

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Oleg-OMON/gin-rest-api.git/models"
	"github.com/Oleg-OMON/gin-rest-api.git/store"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *store.Store
}

func (h *Handler) AllPlayers(c *gin.Context) {
	rows, err := h.DB.DataBase.Query("Select * From players")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	players := []models.Player{}

	for rows.Next() {
		pl := models.Player{}
		err := rows.Scan(&pl.Player_id, &pl.First_name, &pl.Last_name, &pl.Nickname, &pl.Citizenship, &pl.Dob, &pl.Role)
		if err != nil {
			fmt.Println(sql.ErrNoRows)
			continue
		}
		players = append(players, pl)
	}

	c.IndentedJSON(http.StatusOK, players)

}

func (h *Handler) AllGames(c *gin.Context) {
	rows, err := h.DB.DataBase.Query("Select * From games")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	games := []models.Game{}

	for rows.Next() {
		game := models.Game{}
		err := rows.Scan(&game.Game_id, &game.Team, &game.City, &game.Goals, &game.Game_date, &game.Own)
		if err != nil {
			fmt.Println(sql.ErrNoRows)
			continue
		}
		games = append(games, game)
	}

	c.IndentedJSON(http.StatusOK, games)

}
