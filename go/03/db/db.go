package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	dsn := "dbname=docker user=docker password=docker host=localhost sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err.Error())
	}

	return db
}
