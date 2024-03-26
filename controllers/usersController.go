package controllers

import (
	"net/http"

	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/Man4ct/belajar-golang-gorm/models"
	"github.com/gin-gonic/gin"
)

func UsersCreate(c *gin.Context) {

	user := models.User{
		Username: "test",
		Email:    "test@gmail.com",
		Password: "password",
		FullName: "test",
		Role:     "r",
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Limit(10).Order("id asc").Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"count": len(users),
	})
}

func GetUser(c *gin.Context) {
	var user models.User
	result := initializers.DB.First(&user, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	result := initializers.DB.First(&user, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}

	var updateUser models.UserUpdate
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := initializers.DB.Model(&user).Updates(updateUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {
	var user models.User
	result := initializers.DB.Delete(&user, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully.",
	})
}
