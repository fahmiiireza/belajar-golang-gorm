package db

import "gorm.io/gorm"

type Shelf struct {
	gorm.Model

	Floor        int    `gorm:"not null"`
	Section      int    `gorm:"not null"`
	BookCapacity int    `gorm:"not null"`
	Books        []Book `gorm:"foreignKey:ShelfID"`
}
