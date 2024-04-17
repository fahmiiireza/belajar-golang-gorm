package helper

import (
	"regexp"

	model "github.com/Man4ct/belajar-golang-gorm/db/model"
	"gorm.io/gorm"
)

func CreateUser(tx *gorm.DB, username, email, password, fullName string, role model.Role) (model.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		FullName: fullName,
		Role:     role,
	}

	err = tx.Create(&user).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
func IsValidEmail(email string) bool {
	// Regular expression for email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
