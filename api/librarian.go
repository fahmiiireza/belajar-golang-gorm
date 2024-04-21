package api

import (
	"net/http"
	"time"

	"github.com/Man4ct/belajar-golang-gorm/db"
	model "github.com/Man4ct/belajar-golang-gorm/db/model"
	"github.com/Man4ct/belajar-golang-gorm/helper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createLibrarian(c *gin.Context) {
	var librarianRequest LibrarianRequest
	var librarian model.Librarian
	if err := c.BindJSON(&librarianRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if librarianRequest.EmploymentStatus == "RESIGNED" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employment status cannot be resigned"})
		return
	}

	if !helper.IsValidEmail(librarianRequest.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	if exists, err := helper.CheckExistingUser(librarianRequest.Username, librarianRequest.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check for existing user"})
		return
	} else if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with that username or email already exists"})
		return
	}

	usnLoggedIn, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User logged in info not found in context"})
		return
	}

	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		var userIDCreatedBy uint
		if err := tx.Model(&model.User{}).Where("username = ? AND users.deleted_at IS NULL", usnLoggedIn).Select("id").First(&userIDCreatedBy).Error; err != nil {
			return err
		}

		var adminID uint
		if err := tx.Model(&model.Admin{}).Where("user_id = ?", userIDCreatedBy).Select("id").First(&adminID).Error; err != nil {
			return err
		}

		user, err := helper.CreateUser(tx, librarianRequest.Username, librarianRequest.Email, librarianRequest.Password, librarianRequest.FullName, model.RoleLibrarian)
		if err != nil {
			return err
		}

		librarian = model.Librarian{
			Salary:           librarianRequest.Salary,
			EmploymentStatus: librarianRequest.EmploymentStatus,
			JoiningDate:      librarianRequest.JoiningDate,
			CreatedBy:        adminID,
			UserID:           user.ID,
		}
		if err := tx.Create(&librarian).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create librarian"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"librarian": librarian})
}

func getLibrarian(c *gin.Context) {
	var librarian model.Librarian
	result := db.GetDB().
		Preload("User").
		Joins("JOIN users ON users.id = librarians.user_id").
		Where("employment_status != ? AND users.deleted_at IS NULL", "RESIGNED").
		First(&librarian, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"librarian": librarian,
	})
}

func getAllLibrarian(c *gin.Context) {
	var librarians []model.Librarian
	result := db.GetDB().
		Preload("User").
		Joins("JOIN users ON users.id = librarians.user_id").
		Where("employment_status != ? AND users.deleted_at IS NULL", "RESIGNED").Find(&librarians)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"librarians": librarians,
	})
}
func updateLibrarian(c *gin.Context) {
	var updateData LibrarianUpdateRequest
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if updateData.User.Email != "" && !helper.IsValidEmail(updateData.User.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email"})
		return
	}

	var librarian model.Librarian
	result := db.GetDB().Where("employment_status != ?", "RESIGNED").Preload("User").First(&librarian, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Librarian not found"})
		return
	}

	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&librarian.User).Updates(updateData.User).Error; err != nil {
			return err
		}
		if err := tx.Model(&librarian).Updates(updateData.Librarian).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update librarian"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"librarian": librarian})
}

func deleteLibrarian(c *gin.Context) {
	var librarian model.Librarian
	result := db.GetDB().Where("employment_status != ?", "RESIGNED").Preload("User").First(&librarian, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Librarian not found"})
		return
	}

	librarian.EmploymentStatus = model.EmploymentStatusResigned
	librarian.User.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	if err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&librarian).Updates(librarian).Error; err != nil {
			return err
		}
		if err := tx.Model(&librarian.User).Updates(librarian.User).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete librarian"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Librarian deleted successfully"})
}
