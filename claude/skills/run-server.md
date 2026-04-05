# Skill: Run Server

## Description
Build and start the LibraNet backend server.

## Steps
1. Navigate to `backend/` directory
2. Ensure `.env` file exists (copy from `.env.example` if not)
3. Run `go build -o bin/server ./cmd/server/`
4. Execute `./bin/server`

## Environment Requirements
- PostgreSQL running on configured host:port
- Database created: `createdb libranet`
- `.env` file with valid DB credentials

## Quick Start
```bash
cd backend
cp .env.example .env  # Edit with your DB credentials
make run
```

## Health Check
After server starts, verify: `curl http://localhost:3000/health`
Expected: `{"status":"ok"}`
