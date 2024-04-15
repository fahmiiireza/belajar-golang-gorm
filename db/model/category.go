package db

import "gorm.io/gorm"

type Category struct {
	gorm.Model

	Name        string `gorm:"not null"`
	Description string
	Books       []Book `gorm:"foreignKey:CategoryID"`
}
