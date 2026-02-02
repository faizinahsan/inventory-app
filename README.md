# Inventory Management Application

This repo is for my personal projects for managing inventory app built with **Domain Driven Design (DDD)** architecture in Go.

## 🏗️ Architecture

This application follows **Domain Driven Design (DDD)** principles with Clean Architecture, organized into the following layers:

### Technology Stack

- **Framework**: Go Fiber v2 (High-performance HTTP framework)
- **Configuration**: Viper (Configuration management with .env support)
- **Database**: PostgreSQL with Goose migrations
- **Logging**: Zap (Structured logging)
- **Architecture**: Domain Driven Design with Clean Architecture

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

3. **Start PostgreSQL with Docker**
   ```bash
   # Run PostgreSQL container
   docker run --name postgres \
     -e POSTGRES_USER=user \
     -e POSTGRES_PASSWORD=myAwEsOm3pa55@w0rd \
     -e POSTGRES_DB=inventory_db \
     -p 5432:5432 \
     -d postgres:15
   
   # Or use existing container
   docker start postgres
   ```

4. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration if needed
   # Make sure DNSDB in Makefile matches your Docker setup
   ```

5. **Install development tools (goose migration)**
   ```bash
   make install-tools
   ```

6. **Run database migrations**
   ```bash
   # Check migration status
   make migrate-status
   
   # Run migrations (includes initial schema + dummy data)
   make migrate-up
   ```

   **Note:** The migration includes realistic dummy data with:
   - 5 root categories and 6 subcategories (Electronics, Clothing, etc.)
   - 15 sample products (iPhones, laptops, clothing, etc.)
   - 15 sample transactions (stock in/out/adjustments)
   
   See `DUMMY_DATA.md` for complete details.

7. **Run the application**
   ```bash
   make run
   # or for development with hot reload
   make dev
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

## 🔧 Troubleshooting

### Database Connection Issues

1. **PostgreSQL container not running**
   ```bash
   # Check if container is running
   docker ps | grep postgres
   
   # Start existing container
   docker start postgres
   
   # Or create new container with correct credentials
   docker run --name postgres \
     -e POSTGRES_USER=user \
     -e POSTGRES_PASSWORD=myAwEsOm3pa55@w0rd \
     -e POSTGRES_DB=inventory_db \
     -p 5432:5432 \
     -d postgres:15
   ```

2. **Connection refused**
   ```bash
   # Check if PostgreSQL is accessible
   pg_isready -h localhost -p 5432
   
   # Check Docker container logs
   docker logs postgres
   ```

3. **Database doesn't exist**
   ```bash
   # Create database in existing container
   docker exec -it postgres createdb -U user inventory_db
   
   # Or recreate container with POSTGRES_DB env var
   docker rm -f postgres
   docker run --name postgres \
     -e POSTGRES_USER=user \
     -e POSTGRES_PASSWORD=myAwEsOm3pa55@w0rd \
     -e POSTGRES_DB=inventory_db \
     -p 5432:5432 \
     -d postgres:15
   ```

4. **Migration errors**
   ```bash
   # Check migration status
   make migrate-status
   
   # Try running migrations with --allow-missing flag (already included)
   make migrate-up
   
   # If still failing, check goose installation
   make install-tools
   ```

### Goose Migration Issues

1. **Goose command not found**
   ```bash
   # Install goose tool
   make install-tools
   
   # Or install manually
   GOBIN=/usr/local/bin go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

2. **Migration file format**
   ```bash
   # Create migration with correct format
   make goose-create file=your_migration_name
   
   # This will create: YYYYMMDDHHMMSS_your_migration_name.sql
   ```

3. **Connection string issues**
   ```bash
   # Check DNSDB variable in Makefile matches your Docker setup
   # Default: postgres://user:myAwEsOm3pa55@w0rd@localhost:5432/inventory_db?sslmode=disable
   ```

### Build Issues

1. **Go version compatibility**
   ```bash
   go version  # Should be 1.21 or later
   ```

2. **Dependencies issues**
   ```bash
   go mod download
   go mod tidy
   ```


## 🛠️ Development Tools

### Available Make Commands

```bash
make help              # Show all available commands
make build             # Build the application
make run               # Run the application
make dev               # Run with hot reload (requires air)
make test              # Run tests
make test-coverage     # Run tests with coverage
make clean             # Clean build artifacts
make deps              # Install dependencies
make install-tools     # Install goose migration tool
make goose-create file=migration_name  # Create new migration file
make migrate-up        # Run database migrations
make migrate-down      # Rollback database migrations  
make migrate-status    # Check migration status
make docker-build      # Build Docker image
make docker-run        # Run Docker container
make fmt               # Format code
make lint              # Run linter
```

### Database Migration with Goose

The project uses **goose** for database migrations:

```bash
# Install goose tool
make install-tools

# Create new migration
make goose-create file=add_users_table

# Check migration status
make migrate-status

# Run migrations
make migrate-up

# Rollback last migration
make migrate-down
```

**Migration Files Location:** `internal/infrastructure/database/migrations/`

**Database Connection:** The connection string is defined in Makefile as `DNSDB` variable:
```
postgres://user:myAwEsOm3pa55@w0rd@localhost:5432/inventory_db?sslmode=disable
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
| `DATABASE_HOST` | Database host | `localhost` |
| `DATABASE_PORT` | Database port | `5432` |
| `DATABASE_USER` | Database user | `user` |
| `DATABASE_PASSWORD` | Database password | `myAwEsOm3pa55@w0rd` |
| `DATABASE_DBNAME` | Database name | `inventory_db` |
| `DATABASE_SSLMODE` | Database SSL mode | `disable` |
| `LOGGER_LEVEL` | Log level (debug/info/warn/error) | `info` |
| `LOGGER_FORMAT` | Log format (json/console) | `console` |
| `ENV` | Environment (development/production) | `development` |

**Configuration with Viper:**
- Uses dot notation internally (`server.host`, `database.user`, etc.)
- Environment variables use underscore format (`SERVER_HOST`, `DATABASE_USER`, etc.)
- Supports both .env files and environment variables
- Automatic conversion between formats (e.g., `server.host` ↔ `SERVER_HOST`)

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

