# LibraNet - Claude Code Project Guide

## Project Overview
LibraNet is a full-stack library management system with a Go REST API backend and React frontend. Supports JWT-based authentication and role-based authorization (user/admin).

## Tech Stack

### Backend
- **Language:** Go 1.22+
- **Web Framework:** Fiber v2
- **ORM:** Bun (uptrace/bun) with PostgreSQL
- **Authentication:** JWT (golang-jwt/jwt/v5)
- **Password Hashing:** bcrypt (golang.org/x/crypto)

### Frontend
- **Framework:** React 19 + TypeScript
- **Build Tool:** Vite
- **Styling:** TailwindCSS v4
- **State Management:** Zustand
- **HTTP Client:** Axios
- **Router:** React Router v7

## Architecture

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
├── Makefile
└── .env.example

frontend/src/
├── api/                        # Axios API client per domain (auth, books, reservations, reports)
├── components/
│   ├── layout/                 # Navbar, Layout (with Outlet)
│   └── ui/                     # ProtectedRoute, shared UI
├── pages/
│   ├── auth/                   # LoginPage, RegisterPage
│   ├── books/                  # BooksPage (search/list), BookDetailPage
│   ├── reservations/           # ReservationsPage, ReservationDetailPage
│   └── admin/                  # AdminBooksPage (CRUD), AdminReportsPage
├── store/authStore.ts          # Zustand auth state (token, user, localStorage)
├── types/index.ts              # TypeScript interfaces matching backend DTOs
└── App.tsx                     # Router with lazy-loaded pages
```

## API Endpoints (12 total)
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

## Frontend Pages (10 total)
| Route               | Component             | Auth       |
|---------------------|-----------------------|------------|
| /login              | LoginPage             | Public     |
| /register           | RegisterPage          | Public     |
| /books              | BooksPage             | Protected  |
| /books/:id          | BookDetailPage        | Protected  |
| /reservations       | ReservationsPage      | Protected  |
| /reservations/:id   | ReservationDetailPage | Protected  |
| /admin/books        | AdminBooksPage        | Admin Only |
| /admin/reports      | AdminReportsPage      | Admin Only |
| /403                | ForbiddenPage         | Public     |
| /404                | NotFoundPage          | Public     |

## Key Conventions

### Backend
- Files: `snake_case.go`, Packages: lowercase, Structs: `PascalCase`
- JSON fields: `camelCase` (matching API spec: `_id`, `userId`, `bookId`)
- All errors return `{"message": "..."}` format
- Service layer defines domain errors as `var`, handlers map to HTTP status codes
- Login lockout: 5 failed attempts = 15 min account lock

### Frontend
- Lazy-loaded route components for code splitting
- Zustand store persists auth to localStorage
- Axios interceptor auto-attaches JWT and handles 401 logout
- 400ms debounce on search input
- Skeleton screens during loading
- aria-label on all interactive elements

## Commands
```bash
# Backend
cd backend && make dev          # Run Go server (port 3000)
cd backend && make build        # Build binary
cd backend && make test         # Run tests

# Frontend
cd frontend && npm run dev      # Dev server (port 5173, proxies /api to :3000)
cd frontend && npm run build    # Production build
cd frontend && npx tsc --noEmit # Type check
```

## Claude Code Resources
- `claude/skills/` - test-backend, run-server, db-migrate
- `claude/hooks/` - pre-commit, post-edit, architecture-guide
- `claude/tasks/` - Task tracking and progress
