package db

import (
	"gorm.io/gorm"
)

type User struct {
	Username string `gorm:"unique;constraint:users_username_key" json:"username"`
	Email    string `gorm:"index:email_index,unique;not null" json:"email"`
	Password string `gorm:"unique" json:"password"`
	FullName string `gorm:"not null" json:"full_name"`
	Role     Role   `gorm:"type:role;not null" json:"role"`
	gorm.Model
}

// type Category struct {
// 	gorm.Model

// 	Name        string `gorm:"not null"`
// 	Description string
// }

// type Book struct {
// 	gorm.Model

// 	ISBN          uint   `gorm:"unique;not null"`
// 	Title         string `gorm:"not null"`
// 	Language      string `gorm:"not null"`
// 	TotalCopy     int    `gorm:"not null"`
// 	CreatedBy     uint   `gorm:"not null"`
// 	ShelfID       uint
// 	CategoryID    uint
// 	CreatedByUser User     `gorm:"foreignKey:CreatedBy"`
// 	Shelf         Shelf    `gorm:"foreignKey:ShelfID"`
// 	Category      Category `gorm:"foreignKey:CategoryID"`
// }

// type Shelf struct {
// 	gorm.Model

// 	Floor        int `gorm:"not null"`
// 	Section      int `gorm:"not null"`
// 	BookCapacity int `gorm:"not null"`
// }

// type Borrow struct {
// 	gorm.Model

// 	BorrowDate   time.Time    `gorm:"not null"`
// 	StudentID    uint         `gorm:"not null"`
// 	LibrarianID  uint         `gorm:"not null"`
// 	BorrowStatus BorrowStatus `gorm:"type:borrow_status;not null"`
// 	DeadlineDate time.Time    `gorm:"not null"`
// 	Librarian    Librarian    `gorm:"foreignKey:LibrarianID"`
// 	Student      Student      `gorm:"foreignKey:StudentID"`
// }

// type BorrowItem struct {
// 	gorm.Model

// 	BookID     uint `gorm:"not null"`
// 	BorrowID   uint `gorm:"not null"`
// 	ReturnDate time.Time
// 	FineAmount int    `gorm:"not null"`
// 	Borrow     Borrow `gorm:"foreignKey:BorrowID"`
// 	Book       Book   `gorm:"foreignKey:BookID"`
// }

// type Class struct {
// 	gorm.Model

// 	Name            string `gorm:"not null"`
// 	StudentCapacity int    `gorm:"not null"`
// }

// type AuthorBook struct {
// 	ID       uint   `gorm:"primaryKey"`
// 	AuthorID uint   `gorm:"not null"`
// 	BookID   uint   `gorm:"not null"`
// 	Author   Author `gorm:"foreignKey:AuthorID"`
// 	Book     Book   `gorm:"foreignKey:BookID"`
// }
