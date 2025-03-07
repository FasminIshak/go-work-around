# Go PostgreSQL Project

This is a basic Go project setup with PostgreSQL database connection.

## Prerequisites

- Go 1.16 or later
- PostgreSQL installed and running
- `github.com/lib/pq` package (installed automatically via go modules)

## Setup

1. Clone the repository
2. Update the database configuration in `config/config.go` with your PostgreSQL credentials
3. Create a database named `test_db` in PostgreSQL (or update the config to use a different database name)
4. Run the application:
   ```bash
   go run main.go
   ```

## Project Structure

- `config/` - Configuration settings
- `db/` - Database connection and query handlers
- `main.go` - Main application entry point

## Database Configuration

The database configuration can be modified in `config/config.go`. Update the following values according to your PostgreSQL setup:

- `DBHost` - Database host (default: "localhost")
- `DBPort` - Database port (default: "5432")
- `DBUser` - Database user (default: "postgres")
- `DBPassword` - Database password
- `DBName` - Database name (default: "test_db") 
