package api

import (
	"net/http"

	"github.com/Man4ct/belajar-golang-gorm/db"
	model "github.com/Man4ct/belajar-golang-gorm/db/model"
	"github.com/Man4ct/belajar-golang-gorm/helper"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	type UserLogin struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbUser model.User
	if err := db.GetDB().Where("username = ?", user.Username).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !helper.CheckPasswordHash(dbUser.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	// Generate JWT token
	tokenString, err := helper.CreateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// // CreateUser creates a new user
// func createUser(c *gin.Context) {
// 	var newUser struct {
// 		Username string     `json:"username" binding:"required"`
// 		Email    string     `json:"email" binding:"required,email"`
// 		Password string     `json:"password" binding:"required"`
// 		FullName string     `json:"full_name" binding:"required"`
// 		Role     model.Role `json:"role" binding:"required"`
// 	}

// 	if err := c.BindJSON(&newUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	hashedPassword, err := helper.HashPassword(newUser.Password)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	user := model.User{
// 		Username: newUser.Username,
// 		Email:    newUser.Email,
// 		Password: hashedPassword,
// 		FullName: newUser.FullName,
// 		Role:     newUser.Role,
// 	}

// 	result := db.GetDB().Create(&user)
// 	if result.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"user": user})
// }

// func getUsers(c *gin.Context) {
// 	var users []model.User
// 	result := db.GetDB().Limit(10).Order("id asc").Find(&users)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": result.Error,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"users": users,
// 		"count": len(users),
// 	})
// }

// func getUser(c *gin.Context) {
// 	var user model.User
// 	result := db.GetDB().First(&user, c.Param("id"))
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": result.Error,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"user": user,
// 	})
// }

// func updateUser(c *gin.Context) {
// 	type UserUpdate struct {
// 		Username string     `json:"username"`
// 		Email    string     `json:"email"`
// 		FullName string     `json:"full_name"`
// 		Role     model.Role `json:"role"`
// 	}

// 	var user model.User

// 	result := db.GetDB().First(&user, c.Param("id"))
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": result.Error,
// 		})
// 		return
// 	}

// 	var updateUser UserUpdate
// 	if err := c.BindJSON(&updateUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := db.GetDB().Model(&user).Updates(updateUser).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)

// }

// func deleteUser(c *gin.Context) {
// 	var user model.User
// 	result := db.GetDB().Delete(&user, c.Param("id"))
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": result.Error,
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "User deleted successfully.",
// 	})
// }
