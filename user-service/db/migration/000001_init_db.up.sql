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
