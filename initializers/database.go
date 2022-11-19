package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DSN")
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error loading database")
	}
}
