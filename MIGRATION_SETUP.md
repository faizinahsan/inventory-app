# Database Migration Setup - Updated

## Overview
Migration setup telah diperbaiki untuk menggunakan goose dengan connection string yang konsisten.

## Key Changes Made

### 1. Makefile Updates
- **DNSDB Variable**: `postgres://user:myAwEsOm3pa55@w0rd@localhost:5432/inventory_db?sslmode=disable`
- **goose-create**: Command untuk membuat migration file baru
- **migrate-up**: Menggunakan `--allow-missing` flag
- **migrate-status**: Check status migration
- **migrate-down**: Rollback migration

### 2. Database Credentials
- **Username**: `user` (bukan `postgres` atau `inventory_user`)
- **Password**: `myAwEsOm3pa55@w0rd`
- **Database**: `inventory_db`
- **Port**: `5432`

### 3. Docker PostgreSQL Setup
```bash
docker run --name postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=myAwEsOm3pa55@w0rd \
  -e POSTGRES_DB=inventory_db \
  -p 5432:5432 \
  -d postgres:15
```

### 4. Migration Commands
```bash
# Install goose
make install-tools

# Create new migration
make goose-create file=migration_name

# Check status
make migrate-status

# Run migrations
make migrate-up

# Rollback
make migrate-down
```

### 5. Files Updated
- ✅ `Makefile` - Updated migration commands dengan goose
- ✅ `README.md` - Updated documentation
- ✅ `.env.example` - Updated database credentials
- ✅ Migration files - Located in `internal/infrastructure/database/migrations/`

## Ready to Use
The migration system is now properly configured and ready for use with Docker PostgreSQL container.
