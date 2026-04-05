# LibraNet Backend - Task Tracker

## Task Status Legend
- `DONE` - Completed and tested
- `IN_PROGRESS` - Currently being worked on
- `PENDING` - Not started yet
- `BLOCKED` - Waiting on dependency
- `REVIEW` - Waiting for user verification (Postman test)

---

## Phase 1: Project Foundation

### Task 1.1 - Git & Go Module Init
**Status:** DONE
**Description:** Initialize git repo, go module, directory structure, dependencies.
**Files:** go.mod, go.sum, .gitignore, .env.example

### Task 1.2 - Config, Database & Models
**Status:** DONE
**Description:** Config loader, Bun DB connection, User/Book/Reservation models.
**Files:** internal/config/config.go, internal/database/database.go, internal/model/*.go

### Task 1.3 - DTO, Middleware & JWT
**Status:** DONE
**Description:** Request/Response DTOs, JWT helper, auth middleware.
**Files:** internal/dto/*.go, internal/middleware/auth.go, pkg/jwt/jwt.go

### Task 1.4 - Repository Layer
**Status:** DONE
**Description:** User, Book, Reservation repositories with Bun queries.
**Files:** internal/repository/*.go

### Task 1.5 - Service Layer
**Status:** DONE
**Description:** Auth, Book, Reservation, Report business logic.
**Files:** internal/service/*.go

### Task 1.6 - Handler Layer
**Status:** DONE
**Description:** Fiber HTTP handlers for all endpoints.
**Files:** internal/handler/*.go

### Task 1.7 - Router & Main Entry Point
**Status:** DONE
**Description:** Route definitions, Fiber app setup, graceful shutdown.
**Files:** internal/router/router.go, cmd/server/main.go

### Task 1.8 - CLAUDE.md & Claude Structure
**Status:** IN_PROGRESS
**Description:** CLAUDE.md, claude/ skills, hooks, tasks.
**Files:** CLAUDE.md, claude/**

### Task 1.9 - Build Test & Makefile
**Status:** PENDING
**Description:** Verify compilation, create Makefile with common commands.
**Files:** Makefile

---

## Phase 2: Endpoint Implementation & Testing (Upcoming)

### Task 2.1 - Auth Endpoints (Register/Login)
**Status:** PENDING
**Postman Tests:**
- POST /api/auth/register with valid data -> 201
- POST /api/auth/register with duplicate email -> 409
- POST /api/auth/login with valid credentials -> 200 + token
- POST /api/auth/login with wrong password -> 401

### Task 2.2 - Book CRUD Endpoints
**Status:** PENDING
**Postman Tests:**
- POST /api/books (admin) -> 201
- GET /api/books -> 200 + paginated list
- GET /api/books/:id -> 200 + book detail
- PUT /api/books/:id (admin) -> 200
- DELETE /api/books/:id (admin) -> 204

### Task 2.3 - Reservation Endpoints
**Status:** PENDING
**Postman Tests:**
- POST /api/reservations -> 201
- GET /api/reservations/:id -> 200
- DELETE /api/reservations/:id -> 204 (return book)
- GET /api/users/:id/reservations -> 200

### Task 2.4 - Report Endpoint
**Status:** PENDING
**Postman Tests:**
- GET /api/reports (admin) -> 200

---

## How to Use This Tracker
1. Claude checks current task status before starting work
2. Updates status to IN_PROGRESS when starting
3. Writes tests and verifies they pass
4. Updates status to REVIEW and lists Postman test instructions
5. Waits for user approval before committing
6. After approval, commits and updates status to DONE
