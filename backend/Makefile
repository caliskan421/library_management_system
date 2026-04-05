.PHONY: build run test migrate dev clean vet fmt

# Build the server binary
build:
	go build -o bin/server ./cmd/server/

# Run the server
run: build
	./bin/server

# Run all tests
test:
	go test ./... -v -count=1

# Run short tests only (skip integration)
test-short:
	go test ./... -v -short

# Run database migration (starts server briefly)
migrate: build
	./bin/server

# Development mode with go run
dev:
	go run ./cmd/server/

# Clean build artifacts
clean:
	rm -rf bin/

# Go vet
vet:
	go vet ./...

# Format all Go files
fmt:
	gofmt -w .

# Check formatting (CI)
fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "Files need formatting:" && gofmt -l . && exit 1)
