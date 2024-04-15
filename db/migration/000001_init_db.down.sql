
-- Drop foreign key constraints with CASCADE option
ALTER TABLE IF EXISTS "books" DROP CONSTRAINT IF EXISTS "books_shelf_id_fkey" CASCADE;

-- Now you can drop the shelves table
DROP TABLE IF EXISTS "shelves";

-- Drop foreign key constraints with CASCADE option
ALTER TABLE IF EXISTS "borrows" DROP CONSTRAINT IF EXISTS "borrows_librarian_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "borrows" DROP CONSTRAINT IF EXISTS "borrows_student_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "borrow_items" DROP CONSTRAINT IF EXISTS "borrow_items_borrow_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "borrow_items" DROP CONSTRAINT IF EXISTS "borrow_items_book_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "books" DROP CONSTRAINT IF EXISTS "books_category_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "books" DROP CONSTRAINT IF EXISTS "books_created_by_fkey" CASCADE;
ALTER TABLE IF EXISTS "books" DROP CONSTRAINT IF EXISTS "books_shelf_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "author_books" DROP CONSTRAINT IF EXISTS "author_books_author_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "author_books" DROP CONSTRAINT IF EXISTS "author_books_book_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "students" DROP CONSTRAINT IF EXISTS "students_class_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "librarians" DROP CONSTRAINT IF EXISTS "librarians_created_by_fkey" CASCADE;
ALTER TABLE IF EXISTS "librarians" DROP CONSTRAINT IF EXISTS "librarians_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "admins" DROP CONSTRAINT IF EXISTS "admins_id_fkey" CASCADE;
ALTER TABLE IF EXISTS "authors" DROP CONSTRAINT IF EXISTS "authors_id_fkey" CASCADE;

-- Now you can drop the remaining tables
DROP TABLE IF EXISTS "borrows";
DROP TABLE IF EXISTS "borrow_items";
DROP TABLE IF EXISTS "books";
DROP TABLE IF EXISTS "author_books";
DROP TABLE IF EXISTS "students";
DROP TABLE IF EXISTS "librarians";
DROP TABLE IF EXISTS "admins";
DROP TABLE IF EXISTS "authors";
DROP TABLE IF EXISTS "categories";
DROP TABLE IF EXISTS "classes";
DROP TABLE IF EXISTS "users";

-- Drop the types with CASCADE option
DROP TYPE IF EXISTS "borrow_status" CASCADE;
DROP TYPE IF EXISTS "sex" CASCADE;
DROP TYPE IF EXISTS "employment_status" CASCADE;
DROP TYPE IF EXISTS "role" CASCADE;
