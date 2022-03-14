package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectTODB() {
	dsn := "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Panic("Failed to connect to database", err.Error())
	}
}
