# Cosmic Snap API

<p align="center">
  <img src="./assets/cosmic-snap.jpeg" alt="Cosmic Snap">
</p>

This project is a marvel snap API to manage cards built using Golang. The API is powered by Go-Gin as a router, GoORM for database communication, SQLite as the database, and Swagger for API documentation.

---

## Features

- Introduction to Golang and building modern APIs
- Development environment setup for creating the API
- Using Go-Gin as a router for route management
- Implementing SQLite as the database for the API
- Using GoORM for communication with the database
- Integrating Swagger for API documentation and testing
- Creating a modern package structure for organizing the project
- Implementing a complete card manage API with endpoints for searching, creating, editing, and deleting cards

## Live Version

You can check the documentation and perform tests to the API by visiting the live version hosted at [here](https://cosmic-snap-production.up.railway.app/swagger/index.html).

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/rodrigofmeneses/cosmic-snap.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Makefile Commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application without generating API documentation.
- `make run-with-docs`: Generate the API documentation using Swag, then run the application.
- `make build`: Build the application and create an executable file named `cosmic-snap`.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove the `cosmic-snap` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing

## Usage

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting cards. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=8080.

## Credits

This project was inspired by [arthur404dev](https://github.com/arthur404dev).

Original [video](https://youtu.be/wyEYpX5U4Vg).
