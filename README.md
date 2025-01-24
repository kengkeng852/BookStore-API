# Bookstore API

## Project Description

The Bookstore API is a RESTful service designed to manage a collection of books. It provides functionality to create, retrieve, update, and delete books in a database. Built using Golang, the API is integrated with a MySQL database for data persistence and utilizes the Gorilla Mux router for handling HTTP routes.

## Features

- Retrieve all books in the collection.
- Fetch details of a specific book by ID.
- Add a new book to the collection.
- Update details of an existing book.
- Remove a book from the collection by ID.



## Set Up the Database

   - Install and configure MySQL.
   - Create a database named `kengdb`.
   - Update the `Connect` function in `app.go` with your MySQL credentials:
     ```go
     database, err := gorm.Open("mysql", "<username>:<password>@tcp(127.0.0.1:3306)/kengdb?charset=utf8&parseTime=True&loc=Local")
     ```


