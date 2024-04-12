package controllers

import (
	"net/http"

	"github.com/Man4ct/belajar-golang-gorm/helpers"
	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/Man4ct/belajar-golang-gorm/models"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	type UserLogin struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query the database for the user with the provided username
	var dbUser models.User
	if err := initializers.DB.Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Compare the hashed password from the database with the plaintext password
	if !helpers.CheckPasswordHash(dbUser.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	tokenString, err := helpers.CreateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	var newUser struct {
		Username string      `json:"username" binding:"required"`
		Email    string      `json:"email" binding:"required,email"`
		Password string      `json:"password" binding:"required"`
		FullName string      `json:"full_name" binding:"required"`
		Role     models.Role `json:"role" binding:"required"`
	}

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := helpers.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user := models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: hashedPassword,
		FullName: newUser.FullName,
		Role:     newUser.Role,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
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
	type UserUpdate struct {
		Username string      `json:"username"`
		Email    string      `json:"email"`
		FullName string      `json:"full_name"`
		Role     models.Role `json:"role"`
	}

	var user models.User

	result := initializers.DB.First(&user, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}

	var updateUser UserUpdate
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
