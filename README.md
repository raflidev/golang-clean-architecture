# Golang Clean Architecture

🚀 **Golang Clean Architecture** is a simple example of implementing _clean architecture_ using Go (Golang), with a well-organized code structure designed to be scalable, maintainable, and easy to develop.

## Features

- ✨ Clean architecture structure (separation of concerns)
- 📦 Modular and easily extensible
- 🛠️ Uses `Gin` for the HTTP server
- 🗄️ Integrated with the database using `Sqlx`

### Project Structure

```bash
├── cmd/
│   └── api/                # Application entry point for API server
├── internal/
│   ├── event/
│   │   ├── delivery/
│   │   │   └── http/        # HTTP handlers and routing
│   │   ├── repository/      # Event repository implementations
│   │   └── usecase/         # Business logic for events
│   ├── models/              # Domain models/entities
│   └── server/              # Server setup and initialization
├── migrations/              # Database migration files (SQL scripts)
├── pkg/
│   └── db/
│       └── postgres/        # PostgreSQL database connection utilities
└── go.mod                   # Go module definition
```

## Prerequisites

- Go 1.21+
- Database (e.g., PostgreSQL)

## Installation

```bash
git clone https://github.com/raflidev/golang-clean-architecture.git
cd golang-clean-architecture
go mod tidy
```

## Running the App

```bash
go run ./cmd/api
```

The application will run on `localhost:4000` by default.
