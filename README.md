# Expensix

An expense manager app for my family. Supports entering data manually and uploading a receipt.

## Features

- **Modern Database Setup**
  - PostgreSQL with ULID (Universally Unique Lexicographically Sortable Identifier) support
  - Database migrations using pure SQL
  - Type-safe database operations with SQLc
  - Prepared statements for security

- **Clean Architecture**
  - Separation of concerns with store/server pattern
  - Dependency injection ready
  - Interface-based design for better testing
  - Modular and extensible structure

- **Developer Experience**
  - Hot reloading support
  - Makefile for common operations
  - Structured logging
  - Environment-based configuration

## Project Structure

```table
expensix/
├── cmd/                    # Application entrypoints
│   └──main.go          # Main application
├── server/                 # HTTP server implementation
├── sqlx/                  # Database utilities and migrations
│   ├── migration/        # SQL migrations
│   └── query/           # SQLc queries
└── tmp/                   # Temporary files (gitignored)
```

## Prerequisites

- Go 1.23 or higher
- PostgreSQL 14 or higher
- SQLc
- Make
- golang-migrate
- templ

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/template.git
   cd expensix
   ```

2. Rename module if you want

3. Set up your environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Install dependencies:

   ```bash
   go mod download
   ```

5. Run database migrations:

   ```bash
   make migrate-up
   ```

6. Generate SQLc code:

   ```bash
   make sqlc
   ```

7. Start the server:

   ```bash
   make dev
   ```

## Database Migrations

- Create a new migration

  ```bash
  make generate_migration
  ```

- Apply migrations:

  ```bash
  make migrateup
  ```

- Rollback migrations:

  ```bash
  make migratedown
  ```

## Development

- Run tests:

  ```bash
  make test
  ```

- Format code:

  ```bash
  make fmt
  ```

- Generate SQLc code:

  ```bash
  make sqlc
  ```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
