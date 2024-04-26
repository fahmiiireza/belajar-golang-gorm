/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.up = function (knex) {
  return knex.raw(
    `
    CREATE TABLE "books" (
        "id" bigserial PRIMARY KEY,
        "isbn" varchar UNIQUE NOT NULL,
        "title" varchar NOT NULL,
        "language" varchar NOT NULL,
        "total_copy" integer NOT NULL,
        "shelf_id" integer,
        "category_id" integer,
        "description" varchar,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
      );
      
      CREATE TABLE "shelves" (
        "id" bigserial PRIMARY KEY,
        "floor" integer NOT NULL,
        "section" integer NOT NULL,
        "book_capacity" integer NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
      );
      
      CREATE TABLE "categories" (
        "id" bigserial PRIMARY KEY,
        "name" varchar NOT NULL,
        "description" varchar,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "updated_at" timestamptz NOT NULL DEFAULT (now()),
        "deleted_at" timestamptz
      );
      CREATE TABLE "authors" (
        "id" bigserial PRIMARY KEY,
        "biography" varchar,
        "nationality" varchar NOT NULL
      );
      
      CREATE TABLE "author_books" (
        "id" bigserial PRIMARY KEY,
        "author_id" integer NOT NULL,
        "book_id" integer NOT NULL
      );
      
      CREATE INDEX "book_shelf_index" ON "books" ("shelf_id");
      CREATE INDEX "book_category_index" ON "books" ("category_id");







ALTER TABLE "author_books" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "author_books" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("shelf_id") REFERENCES "shelves" ("id");


ALTER TABLE "books" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
      
  `
  );
};

/**
 * @param { import("knex").Knex } knex
 * @returns { Promise<void> }
 */
exports.down = function (knex) {
  return knex.raw(
    `
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
    `
  );
};
