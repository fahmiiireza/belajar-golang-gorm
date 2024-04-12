package helpers

import (
	"github.com/Man4ct/belajar-golang-gorm/models"
	"gorm.io/gorm"
)

func CreateUser(tx *gorm.DB, username, email, password, fullName string) (models.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		FullName: fullName,
		Role:     models.RoleLibrarian,
	}

	err = tx.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
