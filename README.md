# Go MySQL CRUD REST API

This project is a CRUD REST API example built with Go and MySQL. The API uses a MySQL database running in Docker and leverages the GORM ORM library for database operations within a layered architecture.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Setup](#setup)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Project Structure](#project-structure)

## Features

- CRUD operations for users (Create, Read, Update, Delete)
- MySQL integration using Docker
- ORM support with GORM
- Soft delete functionality

## Technologies Used

- **Go**: Programming language for backend development
- **MySQL**: Database management
- **Docker**: Containerization for database
- **GORM**: ORM (Object-Relational Mapping) for Go
- **Gin**: Web framework for REST API development

## Setup

### 1. Prerequisites

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/)

### 2. Clone the Repository

```bash
git clone https://github.com/yusufborucu/go-mysql-crud.git
cd go-mysql-crud
```

### 3. Set Up Environment Variables

Update .env file in the root directory with the following variables:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=testuser
DB_PASSWORD=testpassword
DB_NAME=testdb
```

### 4. Start MySQL with Docker

Run the following command to start the MySQL container:

```bash
docker-compose up -d
```

### 5. Install Required Go Packages
```bash
go mod tidy
```

### 6. Run the Application
```bash
go run main.go
```

The API should now be running at http://localhost:8080.

## Usage

You can test the API’s core features using the following example cURL commands.

### Create a User

```bash
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "johndoe@example.com"}'
```

### List All Users

```bash
curl http://localhost:8080/users
```

### Retrieve a Specific User

```bash
curl http://localhost:8080/users/1
```

### Update a User

```bash
curl -X PUT http://localhost:8080/users/1 -H "Content-Type: application/json" -d '{"name": "Jane Doe", "email": "janedoe@example.com"}'
```

### Delete a User

```bash
curl -X DELETE http://localhost:8080/users/1
```

## API Documentation

| HTTP Method | Endpoint      | Description                | Request Body                         | Response                    |
|-------------|---------------|----------------------------|--------------------------------------|-----------------------------|
| GET         | `/users`      | List all users             | None                                 | JSON array of users         |
| GET         | `/users/:id`  | Retrieve a specific user   | None                                 | JSON object of user         |
| POST        | `/users`      | Create a new user          | `{ "name": "John", "email": "john@example.com" }` | Confirmation message or error |
| PUT         | `/users/:id`  | Update an existing user    | `{ "name": "Jane", "email": "jane@example.com" }` | Confirmation message or error |
| DELETE      | `/users/:id`  | Delete a user              | None    

## Project Structure

The project follows a layered architecture to separate concerns and make the codebase more maintainable. Below is an overview of the folder structure and its purpose:

```plaintext
go-mysql-crud
│   main.go               # Main application file
│   .env                  # Environment variables
│   docker-compose.yml    # Docker setup for MySQL
├───configs                # Configuration files
│       db.go
├───controllers           # API request handlers
│       user_controller.go
├───models                # Database models
│       user.go
├───repositories          # Database operations
│       user_repository.go
├───services              # Business logic layer
│       user_service.go
└───utils                 # Helper functions
        response.go
```

### Folder Details

- **configs**: Contains database connection settings and initialization code.
- **controllers**: Manages incoming HTTP requests, calling service methods as needed to handle the business logic.
- **models**: Defines data models representing database tables, managed through GORM.
- **repositories**: Implements direct CRUD operations for each model, handling data retrieval and persistence.
- **services**: Contains business logic, orchestrating data flow between the controllers and repositories.
- **utils**: Provides helper functions for common tasks, like formatting JSON responses.

This structure makes it easy to maintain, test, and scale the application.