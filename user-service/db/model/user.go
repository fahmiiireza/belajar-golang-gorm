package db

import (
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"unique;constraint:users_username_key" json:"username"`
	Email    string `gorm:"index:email_index,unique;not null" json:"email"`
	Password string `gorm:"unique" json:"password"`
	FullName string `gorm:"not null" json:"full_name"`
	Role     Role   `gorm:"type:role;not null" json:"role"`
	gorm.Model
}
