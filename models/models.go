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

func (s MyNullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
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
