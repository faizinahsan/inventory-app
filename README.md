# Inventory Management Application

This repo is for my personal projects for managing inventory app built with **Domain Driven Design (DDD)** architecture in Go.

## 🏗️ Architecture

This application follows **Domain Driven Design (DDD)** principles with Clean Architecture, organized into the following layers:

### Project Structure

```
inventory-app/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── application/                # Application Layer
│   │   ├── dto/                    # Data Transfer Objects
│   │   └── usecases/               # Use Cases (Application Services)
│   ├── domain/                     # Domain Layer (Business Logic)
│   │   ├── entities/               # Domain Entities
│   │   ├── repositories/           # Repository Interfaces
│   │   ├── services/               # Domain Services
│   │   └── valueobjects/           # Value Objects
│   ├── infrastructure/             # Infrastructure Layer
│   │   ├── config/                 # Configuration
│   │   ├── database/               # Database Implementation
│   │   └── http/                   # HTTP Server
│   └── interfaces/                 # Interface Layer
│       ├── handlers/               # HTTP Handlers
│       └── middleware/             # HTTP Middleware
├── pkg/                           # Shared Packages
│   ├── logger/                    # Logging utilities
│   └── utils/                     # Common utilities
├── .env.example                   # Environment variables example
├── Dockerfile                     # Docker configuration
├── Makefile                       # Build and development commands
└── README.md                      # This file
```

### Domain Driven Design Layers

#### 1. Domain Layer (`internal/domain/`)
- **Entities**: Core business objects with identity (`Product`, `Category`, `Transaction`)
- **Value Objects**: Immutable objects without identity (`SKU`)
- **Repository Interfaces**: Contracts for data persistence
- **Domain Services**: Business logic that doesn't belong to a single entity

#### 2. Application Layer (`internal/application/`)
- **Use Cases**: Application-specific business logic
- **DTOs**: Data transfer objects for API communication

#### 3. Infrastructure Layer (`internal/infrastructure/`)
- **Database**: PostgreSQL implementation of repositories
- **Config**: Application configuration management
- **HTTP**: HTTP server and routing

#### 4. Interface Layer (`internal/interfaces/`)
- **Handlers**: HTTP request handlers
- **Middleware**: HTTP middleware components

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Make (optional, for using Makefile commands)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd inventory-app
   ```

2. **Install dependencies**
   ```bash
   make deps
   # or
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Set up the database**
   ```bash
   # Create PostgreSQL database
   createdb inventory_db
   
   # Run migrations
   make migrate-up
   ```

5. **Run the application**
   ```bash
   make run
   # or
   go run cmd/api/main.go
   ```

### Development

For development with hot reload:
```bash
# Install air for hot reload
go install github.com/cosmtrek/air@latest

# Run with hot reload
make dev
```

## 📊 Database Schema

### Tables

- **categories**: Product categories with hierarchical support
- **products**: Main product information
- **transactions**: Inventory movement tracking

### Key Features

- UUID primary keys
- Hierarchical categories (parent-child relationships)
- Stock tracking with min/max thresholds
- Transaction history for audit trail
- Automatic timestamps with triggers

## 🔧 API Endpoints

### Products

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/products` | List products with pagination |
| GET | `/api/v1/products/:id` | Get product by ID |
| POST | `/api/v1/products` | Create new product |
| PUT | `/api/v1/products/:id` | Update product |
| DELETE | `/api/v1/products/:id` | Delete product |
| GET | `/api/v1/products/search?q=term` | Search products |
| GET | `/api/v1/products/low-stock` | Get low stock products |

### Health Check

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Application health status |

## 🧪 Testing

Run tests:
```bash
make test

# With coverage
make test-coverage
```

## 🐳 Docker

### Build and run with Docker

```bash
# Build image
make docker-build

# Run container
make docker-run
```

### Using Docker Compose (when available)

```bash
# Start all services
make docker-compose-up

# Stop all services
make docker-compose-down
```

## 🛠️ Development Tools

### Available Make Commands

```bash
make help              # Show all available commands
make build             # Build the application
make run               # Run the application
make dev               # Run with hot reload
make test              # Run tests
make test-coverage     # Run tests with coverage
make clean             # Clean build artifacts
make deps              # Install dependencies
make migrate-up        # Run database migrations
make migrate-down      # Rollback migrations
make migrate-create    # Create new migration
make docker-build      # Build Docker image
make docker-run        # Run Docker container
make fmt               # Format code
make lint              # Run linter
make install-tools     # Install development tools
```

### Code Quality

Format code:
```bash
make fmt
```

Run linter:
```bash
# Install golangci-lint first
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
make lint
```

## 📝 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_HOST` | Server host | `localhost` |
| `SERVER_PORT` | Server port | `8080` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database user | `inventory_user` |
| `DB_PASSWORD` | Database password | `inventory_pass` |
| `DB_NAME` | Database name | `inventory_db` |
| `DB_SSL_MODE` | Database SSL mode | `disable` |
| `LOG_LEVEL` | Log level (debug/info/warn/error) | `info` |
| `LOG_FORMAT` | Log format (json/console) | `console` |
| `ENV` | Environment (development/production) | `development` |

## 📋 TODO

- [ ] Add Category management endpoints
- [ ] Add Transaction/Inventory movement endpoints
- [ ] Add User authentication and authorization
- [ ] Add API documentation with Swagger
- [ ] Add unit tests
- [ ] Add integration tests
- [ ] Add Docker Compose configuration
- [ ] Add CI/CD pipeline
- [ ] Add monitoring and metrics
- [ ] Add caching layer

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

