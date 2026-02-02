package repositories

import (
	"context"
	"github.com/google/uuid"
	"inventory-app/internal/domain/entities"
)

// ProductRepository defines the interface for product persistence operations
type ProductRepository interface {
	Create(ctx context.Context, product *entities.Product) error
	GetByID(ctx context.Context, id uuid.UUID) (*entities.Product, error)
	GetBySKU(ctx context.Context, sku string) (*entities.Product, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entities.Product, error)
	GetByCategory(ctx context.Context, categoryID uuid.UUID, limit, offset int) ([]*entities.Product, error)
	Update(ctx context.Context, product *entities.Product) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetLowStockProducts(ctx context.Context) ([]*entities.Product, error)
	Search(ctx context.Context, query string, limit, offset int) ([]*entities.Product, error)
}
