package controllers

import (
	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/Man4ct/belajar-golang-gorm/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	user := models.User{
		Username: "test",
		Email:    "test@gmail.com",
		Password: "password",
		FullName: "test",
		Role:     "r",
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}
