# Library Management System Project Specification Document

## 1. Overview

### 1.1 Purpose
The purpose of this Library Management System (LMS) is to provide Greenfield University with a modern, scalable, and efficient platform to manage library resources, roles, and user interactions. The system leverages a microservices architecture to ensure reliability, scalability, and ease of maintenance.

### 1.2 Goals
- **Centralize Book Management**: Allow librarians to manage book inventory efficiently.
- **Improve User Access and Control**: Offer secure, role-based access to admins, librarians, and general users.
- **Enable Self-Service Capabilities**: Allow students and faculty to check book availability and make reservations.
- **Future-Proof the System**: Use a scalable architecture to accommodate the library's growth and evolving needs.

---

## 2. System Architecture

The system uses a **microservices architecture** comprising two main services:

1. **User-Service (Go)**: Handles user authentication, roles, and permissions.
2. **Book-Service (TypeScript)**: Manages book information, availability, and inventory actions.

Each service has its own **PostgreSQL database** to ensure data isolation and performance. The services communicate through RESTful APIs and are containerized with **Docker Compose**.

### 2.1 High-Level Architecture Diagram
*(Placeholder for an architecture diagram if desired.)*

---

## 3. Functional Requirements

### 3.1 User-Service

- **User Authentication**
  - Provide login capabilities for all users (students, librarians, and admins).
  - JWT tokens will be used for session management and to verify user identity.

- **Role Management**
  - Assign roles as follows:
    - **Admins**: Can create and manage librarian accounts.
    - **Librarians**: Can manage book records.
    - **General Users**: Can view book availability and make reservations.

- **Account Operations**
  - Admins can create, read, update, and delete librarian accounts.

#### API Endpoints
1. **POST** `/login` - Login for all users.
2. **POST** `/admins` - Create a new admin account (Admin only).
3. **GET** `/librarians` - View all librarians (Admin only).
4. **POST** `/librarians` - Add a new librarian (Admin only).
5. **PATCH** `/librarians/:id` - Update librarian details (Admin only).
6. **DELETE** `/librarians/:id` - Remove a librarian (Admin only).

---

### 3.2 Book-Service

- **Book Inventory Management**
  - Enable librarians to create, read, update, and delete book records in the library catalog.
  
- **Book Availability Tracking**
  - Update and maintain real-time information on book availability.

- **Access Control**
  - Allow only librarians to perform CRUD operations on book records.

- **User Access to Book Data**
  - Allow general users to view available books and search by title, author, or genre.

#### API Endpoints
1. **GET** `/books` - Retrieve all books (Public).
2. **GET** `/books/:id` - Retrieve a single book by ID (Public).
3. **POST** `/books` - Create a new book record (Librarian only).
4. **PATCH** `/books/:id` - Update book information (Librarian only).
5. **DELETE** `/books/:id` - Delete a book (Librarian only).

---

## 4. Non-Functional Requirements

- **Scalability**: The system should handle up to 5,000 concurrent users and a growing database of books.
- **Security**: Sensitive information should be encrypted. Role-based access control is enforced, and JWT tokens are used for user sessions.
- **Performance**: Response times for API requests should be under 1 second for 95% of operations.
- **Reliability**: The system should have 99.9% uptime with fault-tolerant mechanisms for each service.
- **Usability**: The system should be accessible on mobile and desktop devices.
- **Maintainability**: Each microservice should be independently deployable and maintainable.

---

## 5. Data Models

### 5.1 User-Service Database (user-db)

- **Users Table**
  - `id`: Integer, Primary Key
  - `username`: String, Unique
  - `password`: String, Hashed
  - `role`: Enum (Admin, Librarian, User)
  
- **Roles Table**
  - `role_id`: Integer, Primary Key
  - `role_name`: Enum, Defines the role types

---

### 5.2 Book-Service Database (book-db)

- **Books Table**
  - `id`: Integer, Primary Key
  - `title`: String, Book title
  - `author`: String, Book author
  - `genre`: String, Book genre
  - `published_date`: Date, Book publication date
  - `availability_status`: Boolean, Indicates if the book is available or borrowed

---

## 6. Security Specifications

- **Authentication**: JWT tokens are used to secure API endpoints.
- **Authorization**: Role-based access control ensures only authorized users can perform certain actions.
- **Encryption**: Passwords are stored in a hashed format in the user database.
- **Environment Variables**: Sensitive information (e.g., DB URLs, SECRET_KEY) is managed through environment variables.

---

## 7. Deployment and Maintenance

### 7.1 Containerization with Docker
The project is containerized with Docker, making deployment and scaling easier. Each service is set up with a **Docker Compose** file for local development and deployment.

### 7.2 Continuous Integration/Continuous Deployment (CI/CD)
A CI/CD pipeline will be configured to automatically test, build, and deploy the services upon code changes. Automated tests for both unit and integration levels will be included.

### 7.3 Logging and Monitoring
- **Logging**: Each service will log requests and errors for auditing and debugging.
- **Monitoring**: A monitoring tool (e.g., Prometheus) will be configured to track service health, response times, and error rates.

---

## 8. Testing Strategy

### 8.1 Unit Testing
- Unit tests will cover all core functions, including authentication, CRUD operations, and role-based restrictions.

### 8.2 Integration Testing
- Tests will validate interactions between the user-service and book-service, as well as database interactions.

### 8.3 QA Test Cases
- Test cases will cover various user flows and edge cases, such as login failures, unauthorized access attempts, and data validation checks.

---

## 9. Future Enhancements

- **Borrowing and Returning System**: Introduce a reservation and check-out/check-in functionality.
- **Notifications**: Notify users when reserved books become available.
- **Recommendation Engine**: Suggest books to users based on reading history.

---

This document provides a thorough specification for developing and deploying the Library Management System. Let me know if further detail is needed!
