package models

import "database/sql"

type Player struct {
	Player_id   string
	First_name  string
	Last_name   string
	Nickname    string
	Citizenship sql.NullString
	Dob         string
	Role        string
}
