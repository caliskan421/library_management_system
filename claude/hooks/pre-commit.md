# Hook: Pre-Commit

## Description
Checks to perform before every commit to maintain code quality.

## Checklist
1. **Compilation Check:** `cd backend && go build ./...` must succeed
2. **Test Pass:** `go test ./... -short` must pass
3. **No Secrets:** Verify `.env` is not staged (`git diff --cached --name-only | grep -v .env`)
4. **Go Vet:** `go vet ./...` must pass
5. **Formatting:** `gofmt -l .` should return empty (all files formatted)

## On Failure
- Fix compilation errors before committing
- Fix failing tests before committing
- Remove any accidentally staged secret files
- Run `gofmt -w .` to auto-format

## Token Optimization
- Only run tests for changed packages: `go test ./internal/handler/... -v`
- Skip full test suite for doc-only changes
