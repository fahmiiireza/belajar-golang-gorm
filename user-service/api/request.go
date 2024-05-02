// librarian_request.go

package api

import (
	"time"

	db "github.com/Man4ct/belajar-golang-gorm/db/model"
)

type LibrarianRequest struct {
	Username         string              `json:"username" binding:"required"`
	Email            string              `json:"email" binding:"required,email"`
	Password         string              `json:"password" binding:"required"`
	FullName         string              `json:"fullName" binding:"required"`
	Salary           int                 `json:"salary" binding:"required"`
	EmploymentStatus db.EmploymentStatus `json:"employmentStatus" binding:"required"`
	JoiningDate      time.Time           `json:"joiningDate" binding:"required"`
}

type AdminRequest struct {
	Username         string              `json:"username" binding:"required"`
	Email            string              `json:"email" binding:"required,email"`
	Password         string              `json:"password" binding:"required"`
	FullName         string              `json:"fullName" binding:"required"`
	Salary           int                 `json:"salary" binding:"required"`
	EmploymentStatus db.EmploymentStatus `json:"employmentStatus" binding:"required"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}

type LibrarianUpdate struct {
	Salary           int                 `json:"salary"`
	EmploymentStatus db.EmploymentStatus `json:"employmentStatus"`
}

type LibrarianUpdateRequest struct {
	User      UserUpdate      `json:"user"`
	Librarian LibrarianUpdate `json:"librarian"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	Language    string `json:"language" binding:"required"`
	TotalCopy   int    `json:"totalCopy" binding:"required"`
	ShelfID     uint   `json:"shelfId" binding:"required"`
	CategoryID  uint   `json:"categoryId" binding:"required"`
	Description string `json:"description"`
}
