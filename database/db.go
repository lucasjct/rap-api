package database

import (
	"log"

	"github.com/lucasjct/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DataBaseConnect() {
	stringConnection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		log.Panic("Connection error")
	}
	DB.AutoMigrate(&models.Artist{}) // automigrate of struct's data.
}
