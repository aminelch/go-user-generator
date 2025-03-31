# go-user-generator API

![License](https://img.shields.io/gitlab/license/aminelch/go-user-generator?color=lightblue&style=for-the-badge)
![GO version](https://img.shields.io/badge/version-1.23.4-blue?logo=go&style=for-the-badge)
![Build Status](https://img.shields.io/gitlab/pipeline/aminelch/go-user-generator/main?label=build&style=for-the-badge)
![Test Status](https://img.shields.io/gitlab/pipeline/aminelch/go-user-generator/main?label=test&style=for-the-badge)
![Lint Status](https://img.shields.io/gitlab/pipeline/aminelch/go-user-generator/main?label=lint&style=for-the-badge)
[![Lint Status](https://img.shields.io/badge/live-demo-darkred?style=for-the-badge)](https://go-user-generator.onrender.com)

A simple REST API built with Golang, Gin framework, and SQLite. This API generates random user data and provides a health check endpoint.

## Features
1. `/health`:  Returns API status information including:
- Current version (from APP_VERSION or default)
- Server uptime (start timestamp in ISO 8601)
- Hostname of the server 
- Current timestamp 
- Status indicator ("ok")

2. `/generate`: Generates and returns a random user from the database.
3. SQLite database with migration and pre-seeded users.
4. Logging support for API requests and errors.

## Prerequisites
- Go (>=1.18)
- Gin (>=v1.10.0)
- SQLite (>=v1.5.7)

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

The API will be available at [http://localhost:8080](http://localhost:8080).

## API Endpoints

### Health Check
```bash
GET /health
```
**Response:**
```json
{
  "status": "ok",
  "uptime": "2025-03-25T09:15:00Z",
  "version": "2.3.0",
  "hostname": "prod-server-42",
  "timestamp": "2025-03-25T10:30:22Z"
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
  "email": "john.doe@example.com",
  "uuid": "a7983f8d-8d77-4bda-bd45-819d7e19290c"
}
```

## License
This project is licensed under the GNU General Public License v3.0.