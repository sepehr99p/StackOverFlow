# Stack Overflow Clone API

A RESTful API for a Stack Overflow clone built with Go, Gin, and GORM.

## Features

- User authentication and authorization
- Question and answer and comment management
- Tag system
- Voting system
- Redis caching
- MySQL database

## Prerequisites

- Docker and Docker Compose
- Go 1.24.1 or later

## Tech Stack

- **Backend**: Go 1.24.1
- **Framework**: Gin
- **ORM**: GORM
- **Database**: MySQL
- **Cache**: Redis
- **Documentation**: Swagger

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```env
DB_HOST=db
DB_USER=dev
DB_PASSWORD=devpass
DB_NAME=stackoverflow
DB_PORT=3306
JWT_SECRET=your_jwt_secret
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASS=devpass
```

## Docker Setup

1. Build and start the containers:

```bash
docker-compose up --build
```

2. To stop the containers:

```bash
docker-compose down
```

3. To view logs:

```bash
docker-compose logs -f
```

## API Documentation

Once the application is running, you can access the Swagger documentation at:

```
http://localhost:8080/swagger/index.html
```

## Project Structure

```
.
├── controllers/     # Request handlers
├── middleware/      # Custom middleware
├── models/          # Database models
├── routes/          # API routes
├── services/        # Business logic
├── utils/           # Utility functions
├── .env            # Environment variables
├── docker-compose.yml
├── Dockerfile
└── main.go         # Application entry point
```

## Development

To run the application locally without Docker:

1. Install dependencies:
```bash
go mod download
```

2. Start the application:
```bash
go run main.go
```

## License

This project is licensed under the MIT License.
