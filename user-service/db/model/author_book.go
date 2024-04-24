package db

import (
	"time"

	"gorm.io/gorm"
)

type AuthorBook struct {
	Author    []Author `gorm:"primaryKey"`
	Book      []Book   `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
