package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	RoleAuthor    Role = "AUTHOR"
	RoleAdmin     Role = "ADMIN"
	RoleLibrarian Role = "LIBRARIAN"
	RoleStudent   Role = "STUDENT"
)

type EmploymentStatus string

const (
	EmploymentStatusFullTime EmploymentStatus = "FULLTIME"
	EmploymentStatusPartTime EmploymentStatus = "PARTTIME"
	EmploymentStatusIntern   EmploymentStatus = "INTERN"
)

type Sex string

const (
	SexMale   Sex = "MALE"
	SexFemale Sex = "FEMALE"
)

// type BorrowStatus string

// const (
// 	BorrowStatusBorrowed      BorrowStatus = "BORROWED"
// 	BorrowStatusPartialReturn BorrowStatus = "PARTIAL_RETURN"
// 	BorrowStatusAllReturn     BorrowStatus = "ALL_RETURN"
// )

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

// Foreign Key Constraints
// func (Admin) TableName() string      { return "admins" }
// func (Librarian) TableName() string  { return "librarians" }
// func (Category) TableName() string   { return "categories" }
// func (Book) TableName() string       { return "books" }
// func (Shelf) TableName() string      { return "shelves" }
// func (Borrow) TableName() string     { return "borrows" }
// func (BorrowItem) TableName() string { return "borrow_items" }
// func (Student) TableName() string    { return "students" }
// func (Class) TableName() string      { return "classes" }
// func (AuthorBook) TableName() string { return "author_books" }
