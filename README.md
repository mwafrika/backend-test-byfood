# Book Management System

## Overview
The Book Management System is a powerful and flexible server application built with Go (Golang) and the Gin framework, designed to efficiently manage books. It provides a RESTful API for performing various book-related operations such as adding, updating, deleting, and retrieving books. Additionally, it offers functionality to process URLs for redirection and canonicalization

## Table of Contents
- [API Documentation]()
- [Features]()
- [Installation]()
- [Project Structure]()
- [Endpoints]()
- [Running Tests]()
- [Screenshots]()
- [Contributing]()
- [License]()

## API Documentation
- Document each endpoint using Swagger to provide an interactive API reference.
- Swagger documentation can be accessed at `/docs` endpoint when the server is running.

## Features
- Add, update, delete, and retrieve books.
- URL processing for redirection and canonicalization.
- Interactive API documentation using Swagger.
- Comprehensive unit tests focusing on different edge cases.
- Smooth local setup for development and testing.

## Installation
To get started with the Book Management System, follow these steps:
1. Clone the repository:

```json git clone https://github.com/mwafrika/book-management-system-backend.git
cd book-management-system-backend


### Endpoints

#### 1. Add Book
- **Method**: POST
- **Endpoint**: `/books`
- **Description**: Add a new book to the system.
- **Request Body**:
  ```json
  {
    "title": "string",
    "author": "string",
    "year": number
  }
