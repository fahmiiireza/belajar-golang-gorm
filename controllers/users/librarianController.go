package controllers

import (
	"net/http"
	"time"

	"github.com/Man4ct/belajar-golang-gorm/helpers"
	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/Man4ct/belajar-golang-gorm/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateLibrarian(c *gin.Context) {
	var newLibrarian struct {
		Username         string                  `json:"username" binding:"required"`
		Email            string                  `json:"email" binding:"required,email"`
		Password         string                  `json:"password" binding:"required"`
		FullName         string                  `json:"full_name" binding:"required"`
		Salary           int                     `json:"salary" binding:"required"`
		EmploymentStatus models.EmploymentStatus `json:"employment_status" binding:"required"`
		JoiningDate      time.Time               `json:"joining_date" binding:"required"`
	}

	if err := c.BindJSON(&newLibrarian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := initializers.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	user, err := helpers.CreateUserWithTx(tx, newLibrarian.Username, newLibrarian.Email, newLibrarian.Password, newLibrarian.FullName)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create librarian"})
		return
	}

	librarian := models.Librarian{
		Salary:           newLibrarian.Salary,
		EmploymentStatus: newLibrarian.EmploymentStatus,
		JoiningDate:      newLibrarian.JoiningDate,
		CreatedBy:        1,
		UserID:           user.ID,
	}

	err = tx.Create(&librarian).Error
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create librarian"})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"librarian": librarian})
}

func GetLibrarian(c *gin.Context) {
	var librarian models.Librarian
	result := initializers.DB.Preload("User").Where("employment_status != ?", "RESIGNED").First(&librarian, c.Param("id"))
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
func UpdateLibrarian(c *gin.Context) {
	type UserUpdate struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
	}
	type LibrarianUpdate struct {
		Salary           int                     `json:"salary"`
		EmploymentStatus models.EmploymentStatus `json:"employment_status"`
	}
	type UpdateData struct {
		User      UserUpdate      `json:"user"`
		Librarian LibrarianUpdate `json:"librarian"`
	}

	var updateData UpdateData
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var librarian models.Librarian
	result := initializers.DB.Preload("User").First(&librarian, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	if err := initializers.DB.Transaction(func(tx *gorm.DB) error {
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

func DeleteLibrarian(c *gin.Context) {
	var librarian models.Librarian
	result := initializers.DB.Preload("User").First(&librarian, c.Param("id"))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	librarian.EmploymentStatus = models.EmploymentStatusPartTime
	librarian.JoiningDate = time.Now()
	librarian.User.DeletedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}

	// Update the librarian and associated user
	if err := initializers.DB.Transaction(func(tx *gorm.DB) error {
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

	c.JSON(http.StatusOK, gin.H{"message": "Librarian deleted successfully"})
}
