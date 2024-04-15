package main

import (
	"fmt"

	api "github.com/Man4ct/belajar-golang-gorm/api"
	"github.com/Man4ct/belajar-golang-gorm/db"
	"github.com/Man4ct/belajar-golang-gorm/initializer"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

func init() {
	initializer.LoadEnvVariables()
}

func main() {
	db.ConnectDB()

	r := api.SetupRouter()

	// Start the server
	if err := r.Run(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
