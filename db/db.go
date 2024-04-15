package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

// func init() {
// 	connectDB()
// }

func ConnectDB() {
	// Connect to the database
	var err error
	dsn := os.Getenv("DB_URL")
	dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *gorm.DB {
	return dbInstance
}
