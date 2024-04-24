package db

type Admin struct {
	ID               uint             `gorm:"primaryKey"`
	Salary           int              `gorm:"not null"`
	EmploymentStatus EmploymentStatus `gorm:"type:employment_status;not null"`
	UserID           uint
	User             User
	Librarians       []Librarian `gorm:"foreignKey:CreatedBy"`
}
