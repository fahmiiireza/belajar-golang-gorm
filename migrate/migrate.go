package main

import (
	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/Man4ct/belajar-golang-gorm/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
