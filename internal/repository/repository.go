package repository

import "database/sql"

type Repository struct {
	DataBase *sql.DB
}

func (s *Repository) Open() error {
	db, err := sql.Open("postgres", "user=postgres password=260616 dbname=go_test sslmode=disable")
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.DataBase = db
	return nil
}

func (s *Repository) Close() {
	s.DataBase.Close()
}
