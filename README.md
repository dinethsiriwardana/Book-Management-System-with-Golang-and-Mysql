# Book Management System with Golang and MySQL

This is a Book Management System based on a tutorial by `freeCodeCamp.org`, implemented with Golang and MySQL. This application allows users to manage books by providing CRUD (Create, Read, Update, Delete) operations through a RESTful API.


[![IMAGE ALT TEXT HERE](https://user-images.githubusercontent.com/91774218/226343585-b8ac5425-d720-4a7b-9888-64b713980296.jpg)](https://youtu.be/jFfo23yIWac?t=1235)

## Features

-   Create a new book
-   Get all books
-   Get book by id
-   Update a book
-   Delete a book

## Installation

To install and run this application, follow these steps:

1.  Clone this repository to your local machine. <br>
    ```
    https://github.com/dinethsiriwardana/Book-Management-System-with-Golang-and-Mysql
    ```
2.  Install MySQL and make sure it is running on your machine.
3.  Create a new database named `bookstore`.
4.  Update the MySQL connection string in the `Connect()` function in the `pkg/config/app.go` file.
5.  Build and run the application by executing the following command from the project root directory:

    ```sh
    go run cmd/main/main.go
    ```

## Usage

Once the application is running, you can perform CRUD operations on books by sending HTTP requests to the following endpoints:

| Endpoint          | Method | Description                               |
|-------------------|--------|-------------------------------------------|
| /api/book     | GET    | Get all books            |
| /api/book/:id | GET    |  Get book by id |
| /api/book     | POST   | Create a new book                     |
| /api/book/:id | PUT    |  Update a book |
| /api/book/:id | DELETE | Delete a book |


You can use tools like Postman or cURL to send HTTP requests to the endpoints.

## Contributors

-   [Dineth Siriwardana](https://github.com/dinethsiriwardana)
