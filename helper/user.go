package helper

import (
	model "github.com/Man4ct/belajar-golang-gorm/db/model"
	"gorm.io/gorm"
)

func CreateUser(tx *gorm.DB, username, email, password, fullName string) (model.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		FullName: fullName,
		Role:     model.RoleLibrarian,
	}

	err = tx.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
