package db

import "gorm.io/gorm"

type Class struct {
	gorm.Model

	Name            string `gorm:"not null"`
	StudentCapacity int    `gorm:"not null"`
	Students        []Student
}
