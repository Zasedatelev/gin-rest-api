package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Oleg-OMON/gin-rest-api.git/internal/models"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/repository"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type GameHandler struct {
	DB *repository.Repository
}

func NewGameHandler(DB *repository.Repository) GameHandler {
	return GameHandler{DB}
}

// GetAllPlayers godoc
// @Summary      Get all players
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {integer} integer 1
// @Router       /api/games/all_players [get]
func (h *GameHandler) GetAllPlayers(c *gin.Context) {
	// TODO: выводит данные о всех игроках
	stmt, err := h.DB.DataBase.Preparex("Select * From players")
	if err != nil {
		panic(err)
	}
	data, err := stmt.Queryx()
	if err != nil {
		log.Debug(err)
	}
	defer data.Close()

	players := []models.Player{}
	for data.Next() {
		pl := models.Player{}
		err := data.Scan(&pl.PlayerId, &pl.FirstName, &pl.LastName, &pl.Nickname, &pl.Citizenship, &pl.Dob, &pl.Role)
		if err != nil {
			log.Debug(sql.ErrNoRows)
			continue
		}
		players = append(players, pl)
	}

	c.IndentedJSON(http.StatusOK, players)

}

// AllGames godoc
// @Summary      Get all games
// @Tags         games
// @Accept       json
// @Produce      json
// @Success      200  {integer} integer 1
// @Router       /api/games/all_games [get]
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

// ResultGames godoc
// @Summary      get games involving the player
// @Description  get list by nicname
// @Tags         games
// @Accept       json
// @Produce      json
// @Param        input  body string  true  "Player nickname"
// @Success      200  {integer} integer 1
// @Router       /api/games/results_games/:nickname [get]
func (h *GameHandler) ResultGames(c *gin.Context) {
	// TODO: выводит об играх футболиста, по его nickname
	nickname := c.Param("nickname")
	rows, err := h.DB.DataBase.Query(`SELECT p.nickname, g.team, c.start, c.time_in, c.goals, c.cards FROM lineups c, players p, games g  WHERE c.player_id = p.player_id AND c.game_id = g.game_id AND p.nickname = $1`, nickname)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	results := []models.ResultModelsPlayerLineup{}

	for rows.Next() {
		r := models.ResultModelsPlayerLineup{}
		err := rows.Scan(&r.Nickname, &r.Team, &r.Start, &r.TimeIn, &r.Goals, &r.Cards)
		if err != nil {
			fmt.Println(sql.ErrNoRows)
			continue
		}
		results = append(results, r)
	}
	log.Info(results)
	c.IndentedJSON(http.StatusOK, results)

}

// GetPlayer godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User Nickname"
// @Success 200 {object} models.Player
// @Router /api/games/get_player/:nickname [get]
func (h *GameHandler) GetPlayer(c *gin.Context) {
	// вывод данных о игроке
	nickname := c.Param("nickname")
	stmt, err := h.DB.DataBase.Preparex(`SELECT * FROM players WHERE nickname = $1`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result := []models.Player{}
	rows, err := stmt.Queryx(nickname)
	for rows.Next() {
		player := models.Player{}
		err = rows.Scan(&player.PlayerId, &player.FirstName, &player.LastName, &player.Nickname, &player.Citizenship, &player.Dob, &player.Role)
		if err != nil {
			log.Debug(sql.ErrNoRows)
		}

		result = append(result, player)
	}

	c.IndentedJSON(http.StatusOK, result)
}
