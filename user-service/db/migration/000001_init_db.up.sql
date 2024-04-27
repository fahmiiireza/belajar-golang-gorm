-- migrate -path db/migration -database "postgresql://postgres:postgres@user-db:5432/user-db?sslmode=disable" -verbose up 

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

CREATE UNIQUE INDEX "username_index" ON "users" ("username");
CREATE UNIQUE INDEX "email_index" ON "users" ("email");

CREATE TABLE "admins" (
  "id" bigserial PRIMARY KEY,
  "salary" integer NOT NULL,
  "employment_status" employment_status NOT NULL,
  "user_id" integer UNIQUE NOT NULL
);

CREATE TABLE "librarians" (
  "id" bigserial PRIMARY KEY,
  "salary" integer NOT NULL,
  "employment_status" employment_status NOT NULL,
  "joining_date" date NOT NULL,
  "user_id" integer UNIQUE NOT NULL
);

ALTER TABLE "admins" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
ALTER TABLE "librarians" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;


-- SEEDING
-- Insert data into the users table
-- INSERT INTO users (username, password, full_name, email, role)
-- VALUES
--     ('admin', 'admin_password', 'Admin User', 'admin@example.com', 'ADMIN'),
--     ('librarian1', 'librarian_password', 'Librarian One', 'librarian1@example.com', 'LIBRARIAN'),
--     ('librarian2', 'librarian_password', 'Librarian Two', 'librarian2@example.com', 'LIBRARIAN'),
--     ('student1', 'student_password', 'Student One', 'student1@example.com', 'STUDENT'),
--     ('student2', 'student_password', 'Student Two', 'student2@example.com', 'STUDENT');

-- -- Insert data into the admins table
-- INSERT INTO admins (salary, employment_status, user_id)
-- VALUES
--     (50000, 'FULLTIME', (SELECT id FROM users WHERE username = 'admin'));

-- -- Insert data into the librarians table
-- INSERT INTO librarians (salary, employment_status, joining_date, user_id)
-- VALUES
--     (40000, 'FULLTIME', '2023-01-01', (SELECT id FROM users WHERE username = 'librarian1')),
--     (35000, 'PARTTIME', '2023-01-01', (SELECT id FROM users WHERE username = 'librarian2'));

