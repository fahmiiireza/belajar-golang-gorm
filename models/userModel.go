package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"index:username_index,unique" json:"username"`
	Email    string `gorm:"index:email_index,unique" json:"email"`
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

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Role     Role   `json:"role"`
}
