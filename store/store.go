package store

import (
	"database/sql"
	"fmt"
)

type Store struct {
	DataBase *sql.DB
}

// func New(s *Store) *Store {
// 	return &Store{}
// }

func (s *Store) OpenConnect() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=260616 dbname=football_database sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}

	s.DataBase = db

}
