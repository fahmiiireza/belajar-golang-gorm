package db

type Author struct {
	ID          uint `gorm:"primaryKey"`
	Biography   string
	Nationality string `gorm:"not null"`
	UserID      uint
	User        User
	Books       []Book `gorm:"many2many:author_books;"`
}
