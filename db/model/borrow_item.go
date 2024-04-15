package db

import (
	"time"

	"gorm.io/gorm"
)

type BorrowItem struct {
	gorm.Model

	ReturnDate time.Time
	FineAmount int  `gorm:"not null"`
	BookID     uint `gorm:"not null"`
	BorrowID   uint `gorm:"not null"`
}
