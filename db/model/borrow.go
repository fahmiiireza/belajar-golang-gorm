package db

import (
	"time"

	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model

	BorrowDate   time.Time    `gorm:"not null"`
	StudentID    uint         `gorm:"not null"`
	LibrarianID  uint         `gorm:"not null"`
	BorrowStatus BorrowStatus `gorm:"type:borrow_status;not null"`
	DeadlineDate time.Time    `gorm:"not null"`
	BorrowItems  []BorrowItem
}
