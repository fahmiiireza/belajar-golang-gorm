package helpers

import (
	"github.com/Man4ct/belajar-golang-gorm/initializers"
	"github.com/Man4ct/belajar-golang-gorm/models"
	"gorm.io/gorm"
)

// CreateUser creates a new user with the provided data and returns it
func CreateUser(newUser struct {
	Username, Email, Password, FullName string
	Role                                models.Role
}) (models.User, error) {
	// Hash the password
	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		return models.User{}, err
	}

	// Create the user object
	user := models.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: hashedPassword,
		FullName: newUser.FullName,
		Role:     newUser.Role,
	}

	// Create the user in the database
	if err := initializers.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CreateUserWithTx(tx *gorm.DB, username, email, password, fullName string) (models.User, error) {
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
