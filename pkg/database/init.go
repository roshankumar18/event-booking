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

// func InitDB() {

// 	// dsn := "host=localhost user=postgres password=mypassword dbname=mydb port=5432 sslmode=disable"
// 	dsn := "postgres://postgres:mypassword@localhost:5432/mydb?sslmode=disable"
// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) // Use = to assign to the package variable
// 	if err != nil {
// 		log.Fatalf("Error connecting to the database: %v", err)
// 	}
// 	if err = DB.AutoMigrate(&models.Booking{}); err != nil {
// 		panic(fmt.Sprintf("Failed to auto-migrate: %v", err))
// 	}
// }
