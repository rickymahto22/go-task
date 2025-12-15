# 🚀 Go Backend Task: User Management API

A robust RESTful API built with **Go (Golang)** and **Fiber** that manages user records and dynamically calculates age based on Date of Birth. This project utilizes **PostgreSQL** for persistence and **SQLC** for type-safe database interactions.

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)
![Fiber](https://img.shields.io/badge/Fiber-v2-black?style=flat&logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=flat&logo=postgresql)

## 📋 Features

* **RESTful Architecture:** Complete CRUD operations for User management.
* **Dynamic Logic:** Automatically calculates user `Age` based on `DOB` (Date of Birth) during fetch.
* **Type-Safe SQL:** Uses **SQLC** to generate Go code from raw SQL queries.
* **High Performance:** Built on **GoFiber**, one of the fastest Go web frameworks.
* **Structured Logging:** Implements **Uber Zap** for production-grade logging.
* **Validation:** Request payload validation using `go-playground/validator`.

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Framework:** Fiber v2
* **Database:** PostgreSQL (Neon Cloud / Local)
* **ORM/Generator:** SQLC
* **Driver:** pgx/v5
* **Config:** Godotenv (.env)

## 📂 Project Structure

```text
go-backend-task/
├── cmd/
│   └── server/
│       └── main.go       # Application Entry Point
├── db/
│   ├── migrations/       # Database Schema (SQL)
│   ├── sqlc/             # Generated Go Code (Do not edit manually)
│   └── query.sql         # SQL Queries for SQLC
├── internal/
│   ├── handler/          # HTTP Controllers
│   ├── middleware/       # Logger & RequestID Middleware
│   ├── models/           # API Request/Response Structs
│   ├── repository/       # Database Data Access Layer
│   ├── routes/           # API Route Definitions
│   └── service/          # Business Logic (Age Calculation)
├── .env                  # Environment Variables (Not committed)
├── go.mod                # Dependencies
└── sqlc.yaml             # SQLC Configuration

```

# Go Backend Task

## ⚙️ Setup & Installation

### Prerequisites
* Go 1.20+
* PostgreSQL Database (Local or Cloud like Neon)

### 1. Clone the Repository
```bash
git clone <https://github.com/rickymahto22/go-task>
cd go-backend-task
```
### 2. Install Dependencies
```bash
go mod tidy
```
## 3️⃣ Database Setup

Run the following SQL in your PostgreSQL database to create the required table:

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  dob DATE NOT NULL
);
```
## 4️⃣ Environment Configuration

Create a file named `.env` in the root directory and add your PostgreSQL database connection string:

```env
DB_SOURCE=postgres://user:password@host:port/dbname?sslmode=require
```
## 5️⃣ Run the Server

Start the application using the following command:

```bash
go run cmd/server/main.go
```
## 🔌 API Endpoints

### 1️⃣ Create User

**POST** `/users`

**Request Body:**
```json
{
  "name": "Ricky",
  "dob": "2000-01-01"
}
```
### 2️⃣ Get User (Calculates Age)

**GET** `/users/:id`

**Response:**
```json
{
  "id": 1,
  "name": "Ricky",
  "dob": "2000-01-01",
  "age": 25
}
```
### 3️⃣ List All Users

**GET** `/users`

Returns a list of all users stored in the database.

### 4️⃣ Update User

**PUT** `/users/:id`

Updates the details of an existing user.

**Request Body:**
```json
{
  "name": "Updated Name",
  "dob": "1999-05-10"
}
```

### 5️⃣ Delete User

**DELETE** `/users/:id`

Deletes the user with the specified ID.

## 🧪 Development Notes

- **SQLC Generation:**  
  If you modify SQL queries in `db/query.sql`, run `sqlc generate` to update the Go code.
  
- **Age Calculation:** 
  The logic for calculating age resides in internal/service/user_service.go.


