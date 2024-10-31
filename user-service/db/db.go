package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbInstance *gorm.DB

func ConnectDB() {
	// Connect to the database
	fmt.Println("testestete")
	fmt.Println("testestete")
	fmt.Println("testestete")
	fmt.Println("testestete")
	fmt.Println("testestete")
	var err error
	dsn := os.Getenv("DB_URL") + "?sslmode=disable"
	dbInstance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	fmt.Println(dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = runMigrations()
	if err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
}

func GetDB() *gorm.DB {
	return dbInstance
}

func runMigrations() error {
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	return err
	// }

	dsn := os.Getenv("DB_URL") + "?sslmode=disable"
	fmt.Printf("Database URL: %s\n", dsn)
	// Create a new instance of golang-migrate
	m, err := migrate.New(
		"file:///app/db/migration", // Path to your migration scripts
		dsn,
	)
	if err != nil {
		return err
	}

	// Apply migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
