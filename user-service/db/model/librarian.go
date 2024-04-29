package db

import "time"

type Librarian struct {
	ID               uint             `gorm:"primaryKey"`
	Salary           int              `gorm:"not null"`
	EmploymentStatus EmploymentStatus `gorm:"type:employment_status;not null"`
	JoiningDate      time.Time        `gorm:"not null"`
	CreatedBy        uint             `gorm:"not null"`
	UserID           uint             `gorm:"not null"`
	User             User
}

// type Librarian struct {
// 	ID               uint              `gorm:"primaryKey"`
// 	Salary           *int              // Changed to pointer
// 	EmploymentStatus *EmploymentStatus // Changed to pointer
// 	JoiningDate      *time.Time        // Changed to pointer
// 	CreatedBy        uint
// 	UserID           uint
// 	User             *User
// 	Borrows          []*Borrow // Changed to slice of pointers
// }

// INSERT INTO books (isbn, title, language, total_copy, description) VALUES
// (9780061120084, , 'English', 10, 'To Kill a Mockingbird is a novel by Harper Lee published in 1960. It was immediately successful, winning the Pulitzer Prize, and has become a classic of modern American literature.'),
// (9780743273565, 'The Great Gatsby', 'English', 8, 'The Great Gatsby is a novel by American writer F. Scott Fitzgerald. Set in the Jazz Age on Long Island, the novel depicts narrator Nick Carraway''s interactions with mysterious millionaire Jay Gatsby and Gatsby''s obsession to reunite with his former lover, Daisy Buchanan.'),
// (9780451524935, '1984', 'English', 12, '1984 is a dystopian social science fiction novel by English novelist George Orwell. It was published on 8 June 1949 by Secker & Warburg as Orwell''s ninth and final book completed in his lifetime.'),
// (9780590353427, 'Harry Potter and the Sorcerer''s Stone', 'English', 15, 'Harry Potter and the Philosopher''s Stone is a fantasy novel written by British author J. K. Rowling. The first novel in the Harry Potter series and Rowling''s debut novel, it follows Harry Potter, a young wizard who discovers his magical heritage on his eleventh birthday.'),
// (9780618260300, 'The Hobbit', 'English', 10,
