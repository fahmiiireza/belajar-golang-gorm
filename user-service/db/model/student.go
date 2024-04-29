package db

import "time"

type Student struct {
	ID        uint      `gorm:"primaryKey"`
	Sex       Sex       `gorm:"type:sex;not null"`
	BirthDate time.Time `gorm:"not null"`
	ClassID   uint      `gorm:"not null"`
	UserID    uint
	User      User
}
