# Go-go-go REST API

A RESTful API application built with Go (Golang) for user management with PostgreSQL database integration.

## Overview

This project demonstrates a modern REST API architecture using Go's Gin framework, with PostgreSQL as the persistent data storage solution. The application provides full CRUD operations for user management and is designed with clean architecture principles.

## Features

- **RESTful API Design** - Clean, standardized HTTP endpoints
- **PostgreSQL Integration** - Robust relational database storage
- **User Management** - Complete CRUD operations (Create, Read, Update, Delete)
- **JSON API** - Standard JSON request/response format
- **Clean Architecture** - Organized codebase with separation of concerns

## Technology Stack

- **Language**: Go (Golang)
- **Web Framework**: Gin
- **Database**: PostgreSQL
- **Architecture**: Clean Architecture pattern

## API Endpoints

### Users
- `GET /users` - Retrieve all users
- `GET /users/:id` - Retrieve a specific user by ID
- `POST /users` - Create a new user
- `PUT /users/:id` - Update a user by ID
- `DELETE /users/:id` - Delete a user by ID

## Project Structure

```
Go-go-go/
├── cmd/
│   ├── server/
│   │   └── main.go                 # Application entry point
├── internal/
│   ├── Endpoints/                  # HTTP endpoints(handlers)
│   │   └── Users
│   └── Domain/
│   │    └── models/                # Data models
│   │       └── Users
│   └── Application/
│   │    └── Use Cases/              # Use cases
│   │       └── Users
└── Infrastucture/
│   │    └── Repository/            # Repository
│   │       └── Users
│   │    └── Components/            # Thrid party integrations
│           └── XYZ
├── go.mod                          # Go modules file
├── go.sum                          # Go dependencies
└── README.md                       # Project documentation
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 12 or higher
- Git

## Installation & Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/RasmusThougaardKristensen/Go-go-go.git
   cd Go-go-go
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up PostgreSQL database**
   ```sql
   CREATE DATABASE go_go_go_db;
   CREATE USER supply_user WITH PASSWORD 'your_password';
   GRANT ALL PRIVILEGES ON DATABASE go_go_go_db TO supply_user;
   ```

4. **Configure environment variables**
   ```bash
   export DB_HOST=localhost
   export DB_PORT=5432
   export DB_USER=supply_user
   export DB_PASSWORD=your_password
   export DB_NAME=go_go_go_db
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Usage Examples

### Create a new user
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com"
  }'
```

### Get all users
```bash
curl http://localhost:8080/users
```

### Update a user
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane.doe@example.com"
  }'
```

### Delete a user
```bash
curl -X DELETE http://localhost:8080/users/1
```

## Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o go-go-go-api main.go
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Rasmus Thougaard Kristensen

Project Link: [https://github.com/RasmusThougaardKristensen/Go-go-go](https://github.com/RasmusThougaardKristensen/Go-go-go)
