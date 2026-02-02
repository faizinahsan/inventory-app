package repositories

import (
	"context"
	"github.com/google/uuid"
	"inventory-app/internal/domain/entities"
	"time"
)

// TransactionRepository defines the interface for transaction persistence operations
type TransactionRepository interface {
	Create(ctx context.Context, transaction *entities.Transaction) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Transaction, error)
	GetByProductID(ctx context.Context, productID uuid.UUID, limit, offset int) ([]*entities.Transaction, error)
	GetByType(ctx context.Context, transactionType string, limit, offset int) ([]*entities.Transaction, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time, limit, offset int) ([]*entities.Transaction, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entities.Transaction, error)
}
