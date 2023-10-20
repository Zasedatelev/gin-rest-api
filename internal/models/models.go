package models

import (
	"database/sql"
	"encoding/json"
	"time"

	"gopkg.in/nullbio/null.v4"
)

type MyNullString struct {
	sql.NullString
}

type MyNullFloat64 struct {
	sql.NullFloat64
}

// убирает поле value: true/false при вывове sql.NullString
func (s MyNullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(`null`), nil
}

// Тоже с Float
func (s MyNullFloat64) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.Float64)
	}
	return []byte(`null`), nil
}

type Player struct {
	Player_id   string       `json: "player_id"`
	First_name  string       `json: "first_name"`
	Last_name   string       `json: "last_name"`
	Nickname    string       `json: "nickname"`
	Citizenship MyNullString `json: "citizenship"`
	Dob         string       `json: "dob"`
	Role        string       `json: "role"`
}

type Game struct {
	Game_id   int          `json: "game_id"`
	Team      string       `json: "team"`
	City      MyNullString `json: "city"`
	Goals     null.Uint    `json: "goals"`
	Game_date time.Time    `json: "game_date"`
	Own       null.Uint    `json: "own"`
}

// участие игрока в матче
type Lineups struct {
	Start    string          `json: start`
	Game_id  int             `json: "game_id"`
	Layer_id int             `json: "layer_id"`
	Time_in  sql.NullFloat64 `json: "time_in"` //  число минут, проведенных игроком на поле; NULL, если не выходил.
	Goals    null.Uint       `json: "goals"`   //  число голов, которые игрок забил в матче; NULL, если не забивал
	Cards    MyNullString    `json: "cards"`
}

type ResultModelsPlayerLineup struct {
	// Тут нужна композиция что бы не дублировать поля, а просто ссылать на их тип?
	Nickname string
	Team     string        `json: "team"`
	Start    string        `json: start`
	Time_in  MyNullFloat64 `json: "time_in"`
	Goals    null.Uint     `json: "goals"`
	Cards    MyNullString  `json: "cards`
}
