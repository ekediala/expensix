# Go Web Server Template

A robust, production-ready Go web server template with built-in PostgreSQL support, ULID-based identifiers, tailwindcss, daisyui, htmx, golang-migrate, and SQLc for type-safe database operations.

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
template/
├── cmd/                    # Application entrypoints
│   └── template/          # Main application
├── server/                 # HTTP server implementation
├── store/                 # Database operations and business logic
├── sqlx/                  # Database utilities and migrations
│   ├── migration/        # SQL migrations
│   └── query/           # SQLc queries
└── tmp/                   # Temporary files (gitignored)
```

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 14 or higher
- SQLc
- Make
- golang-migrate
- templ

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/template.git
   cd template
   ```

2. Set up your environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Run database migrations:

   ```bash
   make migrate-up
   ```

5. Generate SQLc code:

   ```bash
   make sqlc
   ```

6. Start the server:

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
