# Skill: Test Backend

## Description
Run Go backend tests and report results.

## Usage
Invoke when code changes are made to verify correctness before committing.

## Steps
1. Navigate to `backend/` directory
2. Run `go test ./... -v -count=1`
3. Check exit code - 0 means all tests pass
4. If tests fail, read error output and fix issues
5. Re-run until all pass

## Pre-conditions
- Go installed and in PATH
- PostgreSQL running (for integration tests) or test uses mocks
- `.env` or `.env.test` configured

## Token Optimization
- Run specific test files when only one package changed: `go test ./internal/service/... -v`
- Use `-run TestName` to run single test during debugging
- Use `-short` flag to skip long-running integration tests
