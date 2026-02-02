.PHONY: build run test clean docker-build docker-run migrate-up migrate-down

# Variables
APP_NAME=inventory-app
DOCKER_IMAGE=inventory-app:latest
DB_URL=postgres://inventory_user:inventory_pass@localhost:5432/inventory_db?sslmode=disable

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
migrate-up:
	@echo "Running database migrations..."
	@if command -v migrate >/dev/null 2>&1; then \
		migrate -path internal/infrastructure/database/migrations -database "$(DB_URL)" up; \
	else \
		echo "migrate tool not installed. Install with: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-down:
	@echo "Rolling back database migrations..."
	@if command -v migrate >/dev/null 2>&1; then \
		migrate -path internal/infrastructure/database/migrations -database "$(DB_URL)" down; \
	else \
		echo "migrate tool not installed. Install with: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
	fi

migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir internal/infrastructure/database/migrations -seq $$name

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
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/cosmtrek/air@latest

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
	@echo "  migrate-create - Create new migration"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container"
	@echo "  fmt            - Format code"
	@echo "  lint           - Run linter"
	@echo "  install-tools  - Install development tools"
	@echo "  help           - Show this help"
