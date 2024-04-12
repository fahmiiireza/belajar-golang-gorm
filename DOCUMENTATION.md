# Library Management System API

Currently there's only User endpoint and Librarian endpoint as the project is still on process of being build


The Library Management System API provides endpoints to manage users, librarians, and search functionality.

## Getting Started

### Installation

1. Install dependencies:
go mod tidy
2. Install golang-migrate
3. Start the server:
go run .


### Migrations

The API requires a PostgreSQL database. Before starting the server, ensure that you have configured the database connection in the `.env` file. Then, run the following command to perform database migrations:

migrate -path migrate -database "your_DB_URL" -verbose up 


## Endpoints

### Users

#### Create a User
- **Verb:** POST
- **URI:** `/user`
- **HTTP Version:** HTTP/1.1
- **Request Header:**
  - Content-Type: application/json
- **Request Body:**
  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string",
    "full_name": "string",
    "role": (enum of models.role can see in models/userModel.go)
  }
  ```

#### Get All Users
- **Verb:** GET
- **URI:** `/users`
- **HTTP Version:** HTTP/1.1

#### Get User by ID
- **Verb:** GET
- **URI:** `/user/:id`
- **HTTP Version:** HTTP/1.1
- **Request Header:**
  - Content-Type: application/json
- **Parameters:**
  - `id` (integer, path, required): The ID of the user.

#### Update User
- **Verb:** PATCH
- **URI:** `/user/:id`
- **HTTP Version:** HTTP/1.1
- **Request Header:**
  - Content-Type: application/json
- **Request Body:**
  ```json
  {
    "username": "string",
    "email": "string",
    "full_name": "string"
  }
  ```
- **Parameters:**
  - `id` (integer, path, required): The ID of the user to update.

#### Delete User
- **Verb:** DELETE
- **URI:** `/user/:id`
- **HTTP Version:** HTTP/1.1
- **Parameters:**
  - `id` (integer, path, required): The ID of the user to delete.

### Librarians

#### Create a Librarian
- **Verb:** POST
- **URI:** `/librarian`
- **HTTP Version:** HTTP/1.1
- **Request Header:**
  - Content-Type: application/json
- **Request Body:**
  ```json
  {
    "username": "string",
    "email": "string",
    "password": "string",
    "full_name": "string",
    "salary": "integer",
    "employment_status": "string"
  }
  ```

#### Get Librarian by ID
- **Verb:** GET
- **URI:** `/librarian/:id`
- **HTTP Version:** HTTP/1.1
- **Request Header:**
  - Content-Type: application/json
- **Parameters:**
  - `id` (integer, path, required): The ID of the librarian.

#### Update Librarian
- **Verb:** PATCH
- **URI:** `/librarian/:id`
- **HTTP Version:** HTTP/1.1
- **Request Header:**
  - Content-Type: application/json
- **Request Body:**
  ```json
  {
    "username": "string",
    "email": "string",
    "full_name": "string",
    "salary": "integer",
    "employment_status": (enum of models.employment_status can see in models/userModel.go)
  }
  ```
- **Parameters:**
  - `id` (integer, path, required): The ID of the librarian to update.

### Search

#### Search Books
- **Verb:** GET
- **URI:** `/search`
- **HTTP Version:** HTTP/1.1
- **Request Header:**
  - Content-Type: application/json
- **Parameters:**
  - `query` (string, query, required): The search query.
