# Gotener

Gotener is a URL shortener written in Go.

## Stack

- Go 1.23+
- PostgreSQL 14+
- goose
- air

## Features

- Shorten long URLs into memorable hashes
- Redirect short URLs to original long URLs
- Collision-free hash generation
- PostgreSQL for persistent storage

## Installation

### Prerequisites

1. Install Go 1.23 or higher
2. Install PostgreSQL 14 or higher

### Setup

1. Clone the repository:

```bash
git clone https://github.com/jonath-z/gotener.git
cd gotener
```

2. Create your environment file:

```bash
cp .env.sample .env
```

3. Configure your .env file:

```bash
PG_USER=postgres
PG_PASSWORD=postgres
PG_HOST=localhost
PG_DATABASE=url_shortener

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres:postgres@localhost:5432/url_shortener
GOOSE_MIGRATION_DIR=./db/migrations
```

4. Run database migrations:

```bash
make database-migrate
```

5. Start the server

```bash
air
```

### API Documentation

**Shorten URL**

- POST `/shortener`
- Body:

```json
{
  "longUrl": "https://example.com/very/long/url"
}
```

- Response:

```json
{
  "longUrl": "https://example.com/very/long/url",
  "shortUrl": "https://short/0a2121c6",
  "hash": "0a2121c6"
}
```

**Redirect to Long URL**

- GET `/{hash}`
- Redirects to the original URL

### Project Structure

```bash
.
├── cmd/           # Application bootstrapping
├── db/            # Database connection and migrations
├── handlers/      # HTTP request handlers
├── lib/           # Shared utilities
├── middleware/    # HTTP middleware
├── models/        # Data models
└── main.go        # Application entry point
```

### Author

Jonathan Z.

- (LinkedIn)[https://www.linkedin.com/in/jonathan-z-0a40ab209/]
