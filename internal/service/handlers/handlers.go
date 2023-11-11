package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Oleg-OMON/gin-rest-api.git/internal/repository"
	"github.com/Oleg-OMON/gin-rest-api.git/internal/service/models"
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
	// TODO: выводит данные о всех игроках
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
	// TODO: выводит данные о всех игроках
	nickname := c.Param("nickname")
	query := `Select p.nickname, g.team, c.start, c.time_in, c.goals, c.cards
			  From lineups c, players p, games g
			  where p.nickname = $1 
			  and c.player_id = p.player_id 
			  and c.game_id = g.game_id`
	rows, err := h.DB.DataBase.Query(query, nickname)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	results := []models.ResultModelsPlayerLineup{}

	for rows.Next() {
		r := models.ResultModelsPlayerLineup{}
		err := rows.Scan(&r.Nickname, &r.Team, &r.Start, &r.Time_in, &r.Goals, &r.Cards)
		if err != nil {
			fmt.Println(sql.ErrNoRows)
			continue
		}
		results = append(results, r)
	}
	c.IndentedJSON(http.StatusOK, results)

}
