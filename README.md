## BookStore Management

This is a basic Bookstore Management System constructed using Golang.

### Features

- Create a new book
- Get all books
- Get a book by ID
- Update a book
- Delete a book

### Prerequisites

- Go 1.20 or later
- Docker

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/to4to/Go-Bookstore-Management.git
    cd go-bookstore-management
    ```

2. Build and run the application using Docker:
    ```sh
    docker build -t go-bookstore-management .
    docker run -p 8080:8080 go-bookstore-management
    ```

### API Endpoints

- `POST /books` - Create a new book
- `GET /books` - Get all books
- `GET /books/{id}` - Get a book by ID
- `PUT /books/{id}` - Update a book
- `DELETE /books/{id}` - Delete a book

### Usage

You can use tools like `curl` or Postman to interact with the API endpoints.

### License

This project is licensed under the MIT License.