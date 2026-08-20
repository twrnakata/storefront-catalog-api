# Storefront Catalog API

REST API for creating and partially updating catalog products.

All responses follow this envelope:

```json
{
  "successful": true,
  "error_code": "",
  "data": null
}
```

Swagger UI is available at `/api-docs` (with both EN/TH specs).

## API Documentation (Swagger)

After the service starts, open:

- Swagger UI: `http://localhost:8080/api-docs`
- English OpenAPI spec: `http://localhost:8080/api-docs/openapi.en.yaml`
- Thai OpenAPI spec: `http://localhost:8080/api-docs/openapi.th.yaml`

## Quick Start

### Prerequisites

- Go `go 1.25.0+` (as in `go.mod`)
- PostgreSQL (via Docker Compose recommended)
- Docker + Docker Compose

### 1) Configure environment

Copy and edit:

```bash
cp .env.example .env
```

For Docker Compose, `docker-compose.yml` uses variables from `.env`.
For local `go run`, the app also reads `.env` (it requires `DATABASE_URL`).

### 2) Start with Docker Compose

```bash
docker compose up -d --build
```

- API: `http://localhost:8080`
- Swagger: `http://localhost:8080/api-docs`

### 3) Start locally (optional)

```bash
go mod download
go run ./cmd/api
```

Then open:
- `http://localhost:8080/api-docs`

## API Endpoints

### POST `/product`

Create product.

Request body (JSON):

```json
{
  "name": "Oolong milk tea",
  "description": "Large cup, less ice",
  "sale_price": 55,
  "price": 65
}
```

Rules:
- `name` and `price` are required
- when provided: `price >= 0`, `sale_price >= 0`
- when both are sent: `sale_price <= price`

Success response example:

```json
{
  "successful": true,
  "error_code": "",
  "data": {
    "data1": "7c4e9b12-3a6f-4d1e-9c82-1f5a8d3e6b07",
    "data2": "Oolong milk tea"
  }
}
```

### PATCH `/product/:id`

Partial update (only fields you send are updated).

Rules (as defined in Swagger/spec + handler validation):
- `id` must be a valid UUID
- body must include at least one field
- `name` cannot be `null`
- `price` cannot be `null`
- `description` and `sale_price` can be `null` to clear the value
- when provided: `price >= 0`, `sale_price >= 0`
- when both are sent: `sale_price <= price`

Success response example:

```json
{
  "successful": true,
  "error_code": "",
  "data": null
}
```

## Project Structure (Clean Architecture + DI)

Key layers:

- `internal/domain/product`: domain interfaces (`ProductRepository`, `ProductService`) and business error (`ErrProductNotFound`)
- `internal/http/handler`: request/response mapping + validation
- `internal/service/product`: orchestrates business logic and maps repository errors to domain errors
- `internal/repository/product`: persistence (GORM) implementation

Dependency wiring happens in:
- `cmd/api/main.go` and `internal/http/route/route.go`

## Testing

### Run all unit tests

```bash
go test ./...
```

### Integration tests (Repository with real Postgres) - “full”

Some repository tests require `DATABASE_URL` to be set.

Recommended flow:
1. Start Postgres:
   ```bash
   docker compose up -d
   ```
2. Ensure your `.env` has `DATABASE_URL` pointing to your Postgres (default in `.env.example` uses `127.0.0.1:5432`).
3. Run:
   ```bash
   go test ./...
   ```

Note: If `DATABASE_URL` is not set, the tests that depend on it will be skipped.

