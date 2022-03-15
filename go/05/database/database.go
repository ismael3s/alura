package database

import (
	"github.com/ismael3s/alura/go/05/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDB() {
	dsn := "host=localhost port=5432 user=student dbname=student password=student sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.Student{})
}
