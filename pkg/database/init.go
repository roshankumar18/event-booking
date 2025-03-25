package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(conn string, model interface{}) {
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database, attempt %d/%d: %v", i+1, maxRetries, err)
		time.Sleep(time.Second * 5) // Wait 5 seconds before retrying
	}
	if err != nil {
		log.Fatalf("Error connecting to the database after %d attempts: %v", maxRetries, err)
	}

	if err = DB.AutoMigrate(model); err != nil {
		panic(fmt.Sprintf("Failed to auto-migrate: %v", err))
	}
}
