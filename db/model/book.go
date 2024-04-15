package db

type Book struct {
	ID          uint      `gorm:"primaryKey"`
	ISBN        string    `gorm:"unique;not null"`
	Title       string    `gorm:"not null"`
	Language    string    `gorm:"not null"`
	TotalCopy   int       `gorm:"not null"`
	Description string    // Add Description field
	CreatedBy   uint      `gorm:"not null"`
	Librarian   Librarian `gorm:"foreignKey:CreatedBy"`
}
