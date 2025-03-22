# go-user-generator API

A simple REST API built with Golang, Gin framework, and SQLite. This API generates random user data and provides a health check endpoint.

## Features
- `/health`: Returns the API version.
- `/generate`: Generates and returns a random user from the database.
- SQLite database with migration and pre-seeded users.
- Logging support for API requests and errors.

## Prerequisites
- Go (>=1.18)
- SQLite

## Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://gitlab.com/aminelch/go-user-generator.git
cd go-user-generator
```

Initialize and download dependencies:

```bash
go mod tidy
```

## Running the API

Build the API:

```bash
make build
```

Run the API:

```bash
make serve
```

The API will be available at (http://localhost:8080)[http://localhost:8080].

## API Endpoints

### Health Check
```bash
GET /health
```
**Response:**
```json
{
  "version": "1.0.0"
}
```

### Generate Random User
```bash
GET /generate
```
**Response Example:**
```json
{
  "id": 3,
  "name": "John Doe",
  "email": "john.doe@example.com"
}
```

## License
This project is licensed under the GNU General Public License v3.0.