package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Oleg-OMON/gin-rest-api.git/internal/models"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/repository"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	DB *repository.Repository
}

func NewGameHandler(DB *repository.Repository) GameHandler {
	return GameHandler{DB}
}

// @Summary      Get all players
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {integer} integer 1
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /games/all_players [get]
func (h *GameHandler) AllPlayers(c *gin.Context) {
	// TODO: выводит данные о всех игроках
	stmt, err := h.DB.DataBase.Preparex("Select * From players")
	if err != nil {
		panic(err)
	}
	data, err := stmt.Queryx()
	if err != nil {
		panic(err)
	}
	defer data.Close()

	players := []models.Player{}

	for data.Next() {
		pl := models.Player{}
		err := data.Scan(&pl.PlayerId, &pl.FirstName, &pl.LastName, &pl.Nickname, &pl.Citizenship, &pl.Dob, &pl.Role)
		if err != nil {
			fmt.Println(sql.ErrNoRows)
			continue
		}
		players = append(players, pl)
	}

	c.IndentedJSON(http.StatusOK, players)

}

// @Summary      Get all games
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {integer} integer 1
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /games/all_games [get]
func (h *GameHandler) AllGames(c *gin.Context) {
	// TODO: выводит данные о всех играх
	stmt, err := h.DB.DataBase.Preparex("Select * From games")
	if err != nil {
		panic(err)
	}
	data, err := stmt.Queryx()
	if err != nil {
		panic(err)
	}
	defer data.Close()

	games := []models.Game{}

	for data.Next() {
		game := models.Game{}
		err := data.Scan(&game.GameId, &game.Team, &game.City, &game.Goals, &game.GameDate, &game.Own)
		if err != nil {
			fmt.Println(sql.ErrNoRows)
			continue
		}
		games = append(games, game)
	}

	c.IndentedJSON(http.StatusOK, games)

}

// @Summary      get games involving the player
// @Description  get list by nicname
// @Tags         games
// @Accept       json
// @Produce      json
// @Param        input  body string  true  "Player nickname"
// @Success      200  {integer} integer 1
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /games/all_players [get]
func (h *GameHandler) ResultGames(c *gin.Context) {
	// TODO: выводит об играх футболиста, по его nickname
	nickname := c.Param("nickname")
	stmt, err := h.DB.DataBase.Preparex(`SELECT p.nickname, g.team, c.start, c.time_in, c.goals, c.cards FROM lineups c, players p, games g  WHERE c.player_id = p.player_id AND c.game_id = g.game_id AND p.nickname = 1$`)

	if err != nil {
		log.Fatal(err)
	}

	data, err := stmt.Queryx(nickname)
	if err != nil {
		panic(err)
	}
	defer data.Close()

	results := []models.ResultModelsPlayerLineup{}

	for data.Next() {
		r := models.ResultModelsPlayerLineup{}
		err := data.StructScan(&r)
		if err != nil {
			fmt.Println(sql.ErrNoRows)
			continue
		}
		results = append(results, r)
	}
	c.IndentedJSON(http.StatusOK, results)

}

func (h *GameHandler) GetPlayer(c *gin.Context) {
	// вывод данных о игроке
	nickname := c.Param("nickname")
	stmt, err := h.DB.DataBase.Preparex("SELECT * FROM players WHERE nickname = $1")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var player models.Player
	err = stmt.Get(&player, nickname)
	if err != nil {
		fmt.Println(sql.ErrNoRows)
	}

	c.IndentedJSON(http.StatusOK, player)
}
