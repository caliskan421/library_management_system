# LibraNet - Claude Code Project Guide

## Project Overview
LibraNet is a library management system REST API built with **Go**, **Fiber** (web framework), and **Bun** (PostgreSQL ORM). It supports JWT-based authentication and role-based authorization (user/admin).

## Tech Stack
- **Language:** Go 1.22+
- **Web Framework:** Fiber v2
- **ORM:** Bun (uptrace/bun) with PostgreSQL
- **Authentication:** JWT (golang-jwt/jwt/v5)
- **Password Hashing:** bcrypt (golang.org/x/crypto)

## Architecture
Layered architecture following Go best practices:

```
backend/
├── cmd/server/main.go          # Entry point, DI, graceful shutdown
├── internal/
│   ├── config/config.go        # Environment configuration
│   ├── database/database.go    # DB connection & migrations
│   ├── model/                  # Bun ORM models (User, Book, Reservation)
│   ├── dto/                    # Request/Response DTOs
│   ├── repository/             # Database queries (Bun)
│   ├── service/                # Business logic
│   ├── handler/                # Fiber HTTP handlers
│   ├── middleware/             # Auth (JWT) & authorization middleware
│   └── router/router.go       # All route definitions
├── pkg/jwt/jwt.go              # JWT token generation/validation
├── tests/                      # Integration & unit tests
├── go.mod / go.sum
├── Makefile
└── .env.example
```

## API Endpoints
| Method | Endpoint                          | Auth     | Role  |
|--------|-----------------------------------|----------|-------|
| POST   | /api/auth/register                | Public   | -     |
| POST   | /api/auth/login                   | Public   | -     |
| GET    | /api/books                        | Required | Any   |
| GET    | /api/books/:bookid                | Required | Any   |
| POST   | /api/books                        | Required | Admin |
| PUT    | /api/books/:bookid                | Required | Admin |
| DELETE | /api/books/:bookid                | Required | Admin |
| POST   | /api/reservations                 | Required | Any   |
| GET    | /api/reservations/:reservationid  | Required | Owner/Admin |
| DELETE | /api/reservations/:reservationid  | Required | Owner/Admin |
| GET    | /api/users/:userid/reservations   | Required | Owner/Admin |
| GET    | /api/reports                      | Required | Admin |

## Key Conventions

### Naming
- Files: `snake_case.go`
- Packages: lowercase single word
- Structs/Types: `PascalCase`
- JSON fields: `camelCase` (matching API spec: `_id`, `userId`, `bookId`)

### Error Handling
- All errors return `{"message": "..."}` format
- Service layer defines domain errors as package-level `var`
- Handlers map service errors to HTTP status codes

### Database
- PostgreSQL with UUID primary keys (gen_random_uuid)
- Bun ORM with struct tags for column mapping
- Automatic table creation via `IfNotExists` migrations

### Authentication Flow
1. Register/Login returns JWT token
2. Protected routes require `Authorization: Bearer <token>` header
3. Middleware extracts userId, userEmail, userRole into Fiber locals
4. Authorize middleware checks role for admin-only routes

## Commands
```bash
make build       # Build binary
make run         # Run server
make test        # Run all tests
make migrate     # Run DB migrations
make dev         # Run with hot reload (if air installed)
```

## Task Tracking
Current tasks are tracked in `claude/tasks/TASKS.md`. Each task has:
- Status: PENDING / IN_PROGRESS / DONE / BLOCKED
- Tests required before completion
- Postman test instructions for user verification

## Claude Code Resources
- `claude/skills/` - Reusable skill definitions
- `claude/hooks/` - Pre/post action hooks
- `claude/tasks/` - Task tracking and progress
