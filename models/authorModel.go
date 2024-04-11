package models

type Author struct {
	ID          uint `gorm:"primaryKey"`
	Biography   string
	Nationality string `gorm:"not null"`
	UserID      uint
	User        User
}
