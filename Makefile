.PHONY: build run test clean docker-build docker-run migrate-up migrate-down

# Variables
APP_NAME=inventory-app
DOCKER_IMAGE=inventory-app:latest
DNSDB=postgres://user:myAwEsOm3pa55@w0rd@localhost:5432/inventory_db?sslmode=disable

# Go commands
build:
	@echo "Building application..."
	go build -o bin/$(APP_NAME) cmd/api/main.go

run:
	@echo "Running application..."
	go run cmd/api/main.go

test:
	@echo "Running tests..."
	go test -v ./...

test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean:
	@echo "Cleaning up..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Development commands
dev:
	@echo "Starting development server with hot reload..."
	air

deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Database commands
# Makefile For Goose
migrate-create:
	goose -dir internal/infrastructure/database/migrations create ${file} sql

migrate-up:
	goose -dir internal/infrastructure/database/migrations postgres "$(DNSDB)" up --allow-missing

migrate-down:
	goose -dir internal/infrastructure/database/migrations postgres "$(DNSDB)" down

migrate-status:
	goose -dir internal/infrastructure/database/migrations postgres "$(DNSDB)" status

# Docker commands
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE)

docker-compose-up:
	@echo "Starting services with docker-compose..."
	docker-compose up -d

docker-compose-down:
	@echo "Stopping services with docker-compose..."
	docker-compose down

# Tools
install-tools:
	@echo "Installing development tools..."
	@echo "Installing goose to /usr/local/bin..."
	@GOBIN=/usr/local/bin go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "goose installed successfully. You can now use 'make migrate-up'"

# Linting and formatting
fmt:
	@echo "Formatting code..."
	go fmt ./...

lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Help
help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  run            - Run the application"
	@echo "  dev            - Run with hot reload (requires air)"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  clean          - Clean build artifacts"
	@echo "  deps           - Install dependencies"
	@echo "  migrate-up     - Run database migrations"
	@echo "  migrate-down   - Rollback database migrations"
	@echo "  migrate-status - Check migration status"
	@echo "  migrate-create - Create new migration"
	@echo "  migrate-simple - Run migration with psql (fallback)"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  fmt            - Format code"
	@echo "  lint           - Run linter"
	@echo "  install-tools  - Install development tools"
	@echo "  help           - Show this help"
