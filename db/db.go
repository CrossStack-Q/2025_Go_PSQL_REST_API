package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgreSQL(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
