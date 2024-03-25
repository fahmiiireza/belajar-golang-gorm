package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Role     Role   `json:"role"`
}
type Role string

const (
	RoleAuthor    Role = "AUTHOR"
	RoleAdmin     Role = "ADMIN"
	RoleLibrarian Role = "LIBRARIAN"
	RoleStudent   Role = "STUDENT"
)
