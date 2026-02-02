package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"inventory-app/internal/infrastructure/config"
)

// DB holds the database connection
type DB struct {
	*sql.DB
}

// NewConnection creates a new database connection
func NewConnection(cfg *config.Config) (*DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL())
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{DB: db}, nil
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.DB.Close()
}

// Health checks if the database is healthy
func (db *DB) Health() error {
	return db.Ping()
}
