-- migrate -path migrate -database "" -verbose up 

-- postgresql://user:password@db:5432/library?sslmode=disable
CREATE TYPE "role" AS ENUM (
  'AUTHOR',
  'ADMIN',
  'LIBRARIAN',
  'STUDENT'
);

CREATE TYPE "employment_status" AS ENUM (
  'FULLTIME',
  'PARTTIME',
  'INTERN',
  'RESIGNED'
);

CREATE TYPE "sex" AS ENUM (
  'MALE',
  'FEMALE'
);

CREATE TYPE "borrow_status" AS ENUM (
  'BORROWED',
  'PARTIAL_RETURN',
  'ALL_RETURN'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz,
  "role" role NOT NULL
);

CREATE TABLE "authors" (
  "id" bigserial PRIMARY KEY,
  "biography" varchar,
  "nationality" varchar NOT NULL,
  "user_id" integer NOT NULL
);

CREATE TABLE "admins" (
  "id" bigserial PRIMARY KEY,
  "salary" integer NOT NULL,
  "employment_status" employment_status NOT NULL,
  "user_id" integer NOT NULL
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

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
  "deleted_at" timestamptz,
  "created_by" integer NOT NULL
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

CREATE TABLE "librarians" (
  "id" bigserial PRIMARY KEY,
  "salary" integer NOT NULL,
  "employment_status" employment_status NOT NULL,
  "joining_date" date NOT NULL,
  "created_by" integer NOT NULL,
  "user_id" integer NOT NULL
);

CREATE TABLE "borrows" (
  "id" bigserial PRIMARY KEY,
  "borrow_date" date NOT NULL,
  "student_id" integer NOT NULL,
  "librarian_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz,
  "borrow_status" borrow_status NOT NULL,
  "deadline_date" date NOT NULL
);

CREATE TABLE "borrow_items" (
  "id" bigserial PRIMARY KEY,
  "book_id" integer NOT NULL,
  "borrow_id" integer NOT NULL,
  "return_date" date,
  "fine_amount" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "students" (
  "id" bigserial PRIMARY KEY,
  "sex" sex NOT NULL,
  "birth_date" date NOT NULL,
  "class_id" integer NOT NULL,
  "user_id" integer NOT NULL
);

CREATE TABLE "classes" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "student_capacity" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE "author_books" (
  "id" bigserial PRIMARY KEY,
  "author_id" integer NOT NULL,
  "book_id" integer NOT NULL
);

CREATE UNIQUE INDEX "username_index" ON "users" ("username");

CREATE UNIQUE INDEX "email_index" ON "users" ("email");

CREATE INDEX "author_nationality_index" ON "authors" ("nationality");

CREATE INDEX "book_shelf_index" ON "books" ("shelf_id");

CREATE INDEX "book_category_index" ON "books" ("category_id");

CREATE INDEX "borrow_student_index" ON "borrows" ("student_id");

CREATE INDEX "borrow_librarian_index" ON "borrows" ("librarian_id");

CREATE INDEX "student_class_index" ON "students" ("class_id");

CREATE INDEX "class_name_index" ON "classes" ("name");

ALTER TABLE "authors" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "admins" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "librarians" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "students" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "students" ADD FOREIGN KEY ("class_id") REFERENCES "classes" ("id") ON DELETE CASCADE;

ALTER TABLE "author_books" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "author_books" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("shelf_id") REFERENCES "shelves" ("id");

ALTER TABLE "borrow_items" ADD FOREIGN KEY ("borrow_id") REFERENCES "borrows" ("id") ON DELETE CASCADE;

ALTER TABLE "borrow_items" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("created_by") REFERENCES "librarians" ("id");

ALTER TABLE "librarians" ADD FOREIGN KEY ("created_by") REFERENCES "admins" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "borrows" ADD FOREIGN KEY ("librarian_id") REFERENCES "librarians" ("id");

ALTER TABLE "borrows" ADD FOREIGN KEY ("student_id") REFERENCES "students" ("id");
