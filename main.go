package main

import (
	"fmt"

	"github.com/Man4ct/belajar-golang-gorm/controllers"
	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository struct {
	*gorm.DB
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {

	r := gin.Default()
	r.POST("/user", controllers.UsersCreate)
	r.GET("/users", controllers.GetUsers)
	r.GET("/user/:id", controllers.GetUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	// Start the server
	if err := r.Run(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
