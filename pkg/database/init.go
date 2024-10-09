package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(conn string, model interface{}) {
	var err error
	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	if err = DB.AutoMigrate(model); err != nil {
		panic(fmt.Sprintf("Failed to auto-migrate: %v", err))
	}
}
