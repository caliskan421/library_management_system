# Hook: Architecture Guide

## Description
Quick reference for navigating the codebase efficiently. Use this to determine where to make changes without reading every file.

## Request Flow
```
HTTP Request
  → Fiber Router (router/router.go)
    → Middleware (middleware/auth.go)
      → Handler (handler/*_handler.go)
        → Service (service/*_service.go)
          → Repository (repository/*_repository.go)
            → Bun ORM → PostgreSQL
```

## Where to Change What

| Task | Primary File | Also Check |
|------|-------------|------------|
| Add new endpoint | router.go + new handler method | middleware if auth needed |
| Change DB schema | model/*.go | repository (queries), dto (response shape) |
| Change business rule | service/*.go | handler (error mapping) |
| Change API response | dto/response.go | handler (JSON return) |
| Change API request | dto/request.go | handler (body parsing + validation) |
| Add auth rule | middleware/auth.go | router.go (middleware placement) |
| Change DB query | repository/*.go | - |

## File Count per Layer (for context budgeting)
- config: 1 file
- database: 1 file
- model: 3 files (user, book, reservation)
- dto: 2 files (request, response)
- repository: 3 files
- service: 4 files (auth, book, reservation, report)
- handler: 4 files
- middleware: 1 file
- router: 1 file
- pkg/jwt: 1 file
- **Total: ~21 Go source files**
