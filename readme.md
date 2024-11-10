# trunct.io

`trunct.io` is a URL shortening service built with Golang, offering features like URL shortening, redirection, and rate limiting. The project is designed with a scalable and maintainable architecture, utilizing `chi` for routing, `pgx` for PostgreSQL interactions, and `jwt` for authentication.

## Features

- **URL Shortening:** Generate unique short codes for any URL.
- **Redirection:** Retrieve and redirect to the original URL using the generated short code.
- **Rate Limiting:** Custom rate limits for different endpoints to protect from abuse.
- **Authentication:** Secure access to user-specific actions using JWT.
- **Data Storage:** PostgreSQL database to manage URL data and user accounts.

## Project Structure

- **`api`**: Manages HTTP routes and middleware setup.
- **`middlewares`**: Custom middleware for authentication and rate limiting.
- **`usecase`**: Contains business logic for handling URLs and user-related operations.
- **`pgstore`**: Database queries and interactions using pgx.

## Technologies Used

- **Go** (1.23.1)
- **sqlc** (Database model and queries creation)
- **tern** (Migration tool)
- **air** (Hot reloading - development)
- **chi** (Router)
- **httprate** (Rate limiting)
- **jwt** (Authentication)
- **pgx** (PostgreSQL driver)
- **Docker** (Containerization)

## Getting Started

### Prerequisites

- **Docker**: Required to run the PostgreSQL database and the Go application in a containerized environment.

### Setup

1.  Clone the repository:

```bash
git clone git@github.com:FelipeBelloDultra/trunct.io.git
```

2. Configure env vars:

```bash
cp .env.example .env
```

3. Run docker containers:

```bash
docker compose up -d
```
