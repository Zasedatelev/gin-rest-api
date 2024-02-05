package repository

import (
	"fmt"

	"github.com/Zasedatelev/gin-rest-api.git/config"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type Repository struct {
	DataBase *sqlx.DB
}

func (s *Repository) Open() error {
	config := config.LoadConfig()

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Postgres.Host, config.Postgres.Port, config.Postgres.User, config.Postgres.Password, config.Postgres.DbName, config.Postgres.SSlMode))
	if err != nil {
		log.WithError(err).Error("ОШИБКА В ПОДКЛБЮЧЕНИИ К БАЗЕ ДАННЫХ")
	}

	if err := db.Ping(); err != nil {
		log.WithError(err).Error("СОЕДИНЕНИЕ НЕ УСТАНОВЛЕНО")
	}
	s.DataBase = db
	return nil
}

func (s *Repository) Close() {
	s.DataBase.Close()
}

// func (s *Repository) AllPlayers() []models.Game {
// 	stmt, err := s.DataBase.Preparex("Select * From players")
// 	if err != nil {
// 		panic(err)
// 	}
// 	data, err := stmt.Queryx()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer data.Close()

// 	games := []models.Game{}

// 	for data.Next() {
// 		game := models.Game{}
// 		err := data.Scan(&game.GameId, &game.Team, &game.City, &game.Goals, &game.GameDate, &game.Own)
// 		if err != nil {
// 			fmt.Println(sql.ErrNoRows)
// 			continue
// 		}
// 		games = append(games, game)
// 	}

// 	return games
// }
