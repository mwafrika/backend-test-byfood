# Library Management System

## Overview
The Library Management System is a powerful and flexible server application built with Go (Golang) and the Gin framework, designed to efficiently manage books. It provides a RESTful API for performing various book-related operations such as adding, updating, deleting, and retrieving books. Additionally, it offers functionality to process URLs for redirection and canonicalization

## Table of Contents
- [API Documentation]()
- [Features]()
- [Installation]()
- [Project Structure]()
- [Endpoints]()
- [Running Tests]()
- [Generating Swagger documentation]()
- [Screenshots]()
- [Contributing]()
- [License]()

## API Documentation
- [Live swagger Documentation](https://reliable-germana-mwafrika-f5c060f5.koyeb.app/api/docs/index.html#/)
- Documented each endpoint using Swagger to provide an interactive API reference.
- Swagger documentation can be accessed at `http://localhost:7000/docs/index.html#/` endpoint when the server is running locally.

## Features
- Add, update, delete, and retrieve books.
- URL processing for redirection and canonicalization.
- Interactive API documentation using Swagger.
- Comprehensive unit tests focusing on different edge cases.
- Smooth local setup for development and testing.

## Installation
To get started with the Book Management System, follow these steps:
1. Clone the repository:

```js
bash git clone https://github.com/mwafrika/byfood-test-backend.git
cd byfood-test-backend 
```

2. Install dependencies:

```js 
go mod tidy
```

3. Setup the development and test database using postgresql.

4. Create a .env file (check the .env.example in the root of the project directory) then provide the values for the following variables in .env file:

```js
PORT=7000
DB_URL=url for your local postgres database
TEST_DB_URL=url for your postgres test database
```
5. Run the local server (CompileDaemon is used for continually running the server in the development environment)

```js
CompileDaemon -command="./byfood-test-backend"
```

## Project Structure
```js
.
├── main.go
├── go.mod
├── go.sum
├── .env
├── README.md
├── controllers
│   ├── book_controller.go
│   └── url_controller.go  
├── models
│   └── book.go
├── services
│   ├── response_formatter_service.go  
│   └── url_service.go
├── tests
│   ├── book_controller_test.go
│   └── url_controller_test.go
├── config
│   ├── database.go
│   ├── loadEnvVariables.go
│   └── logger.go
│   
└── swagger
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
    
```

### Endpoints
The API provides the following endpoints for managing books:

#### 1. Add Book
- **Method**: POST
- **Endpoint**: `POST /api/books`
- **Description**: Add a new book to the system.
- **Request Body**:
 
```js
 {
    "Title": "Going back to the beginning",
    "Author": "Mwafrika",
    "Year": 2003
}
```
- **Response**:
```js
{
  "message": "Book added successfully",
  "data": {
    "id": 1,
    "title": "Going back to the beginning",
    "author": "Mwafrika",
    "Year": 2003
  }
}
```
#### 1. Update a Book
- **Method**: PUT
- **Endpoint**: `PUT /api/books/:id`
- **Description**: Update an existing book.
- **Request Body**:
```js
{
    "title": "Going back to the beginning updated",
    "author": "Mwafrika",
    "Year": 2003,

}
```
- **Response**:
```js
{
  "message": "Book updated successfully",
  "data": {
    "id": 1,
    "title": "Going back to the beginning updated",
    "author": "Mwafrika",
     "created_at": "2024-06-11T11:46:14.39208+03:00",
     "updated_at": "2024-06-11T11:46:14.39208+03:00",
  }
}

```

#### 1. Delete a Book
- **Method**: DELETE
- **Endpoint**: `DELETE /api/books/:id`
- **Description**: Delete a book.
- **Request Body**:

```js
{
  "message": "Book deleted successfully"
}

```
#### 1. Retrieve a Book
- **Method**: GET
- **Endpoint**: `GET /api/books/:id`
- **Description**: Retrieve a book by its ID.
- **Request Body**:
```js
{
  "data": {
    "id": 1,
    "title": "Book Title",
    "author": "Author Name",
     "created_at": "2024-06-11T11:46:14.39208+03:00",
     "updated_at": "2024-06-11T11:46:14.39208+03:00",
  }
}

```

#### 1. List All Books
- **Method**: GET
- **Endpoint**: `GET /api/books`
- **Description**: List all books.
- **Request Body**:
```js
{
    "data": [
        {
            "id": 3,
            "created_at": "2024-06-11T11:46:14.39208+03:00",
            "updated_at": "2024-06-11T11:46:14.39208+03:00",
            "title": "The Great Gatsby",
            "author": "F. Scott Fitzgerald",
            "year": 1925
        },
        {
            "id": 2,
            "created_at": "2024-06-10T16:22:09.660146+03:00",
            "updated_at": "2024-06-10T16:22:09.660146+03:00",
            "title": "Not Well formated book",
            "author": "Mwafrika",
            "year": 202
        }
    ],
    "pagination": {
        "limit": 10,
        "page": 1,
        "total_count": 2
    }
}

```
### Running Tests

To run the tests for the Book Management System, use the following command:

```js
go test ./tests -v 
```

### Generating Swagger documentation
run the following command to generate the swagger documentation then visit the documentation on: 
[http://localhost:7000/api/docs/index.html#/]()

```js
swag init
```

## Screenshots
![](./swaggers.png)

## License
This project is licensed under the MIT License - see the LICENSE file for details.